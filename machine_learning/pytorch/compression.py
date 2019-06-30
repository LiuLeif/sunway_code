from collections import OrderedDict

import matplotlib.pyplot as plt
import numpy as np
import torch
from matplotlib.colors import ListedColormap
from sklearn.datasets import make_moons
from torch.utils.data import DataLoader, Dataset
from tensorboardX import SummaryWriter

writer = SummaryWriter()


# ---------- data ----------
class MoonDataset(Dataset):
    def __init__(self):
        X, Y = make_moons(n_samples=1000, noise=0.2)
        self.X = torch.from_numpy(X).float()
        self.Y = torch.from_numpy(Y).float().view(-1, 1)

    def __getitem__(self, index):
        return self.X[index], self.Y[index]

    def __len__(self):
        return len(self.X)


training_set = MoonDataset()
training_loader = DataLoader(training_set, batch_size=100)

test_set = MoonDataset()
test_loader = DataLoader(test_set, batch_size=1000)


# ---------- helper ----------
def test():
    for x, y in test_loader:
        y_hat = model(x)
        y_hat = y_hat > 0.5
        accu = torch.sum(y_hat == y.byte()).item() / 1000
        return accu


def train():
    for i in range(10):
        for x, y in training_loader:
            loss = criterion(model(x), y)
            optimizer.zero_grad()
            loss.backward()
            optimizer.step()


def visualize():
    plt.close()
    cm = ListedColormap(["#FF0000", "#0000FF"])
    plt.scatter(x=test_set.X[:, 0], y=test_set.X[:, 1], c=test_set.Y[:, 0], cmap=cm)
    xx, yy = np.meshgrid(np.arange(-4, 4, 0.02), np.arange(-4, 4, 0.02))
    X = np.c_[xx.ravel(), yy.ravel()]

    y_hat = model(torch.from_numpy(X).float()).detach().numpy()
    y_hat = (y_hat > 0.5)[:, 0]
    y_hat = y_hat.reshape(xx.shape)

    cm = ListedColormap(["#FF0000", "#0000FF"])
    plt.contour(xx, yy, y_hat, cmap=cm)
    plt.show()


# ---------- quantize ----------
def linear_quantize(input):
    num_bits = 32
    min_q = 0.
    max_q = 2. ** num_bits - 1.
    min_r, max_r = input.min(), input.max()

    # r=S(q-Z)
    #
    # max_val = S (qmax - Z)         max_val - min_val = S (qmax-qmin)
    #                          =>
    # min_val = S (qmin - Z)         Z = qmin-min_val/S

    S = (max_r - min_r) / (max_q - min_q)

    Z = min_q - min_r / S
    Z = np.clip(Z, min_q, max_q)

    Z = int(Z)
    q = Z + input / S
    q.clamp_(min_q, max_q).round_()
    q = q.round()

    ret = S * (q.float() - Z)
    return ret


def quantize():
    state_dict = model.state_dict()
    state_dict_quant = OrderedDict()

    for k, v in state_dict.items():
        state_dict_quant[k] = linear_quantize(v)

    model.load_state_dict(state_dict_quant)
    return model


# ---------- weight pruning ----------
threshold = 1e-1


def compute_mask(weight):
    mask = np.abs(weight) > threshold
    return mask.astype(np.float32)


def pruning():
    # 1. get mask per layer
    # 2. apply mask
    # 3. train
    total_para = 0
    remaining_para = 0

    for layer in model:
        if isinstance(layer, MaskableLinear):
            total_para += layer.weight.nelement()
            mask = compute_mask(layer.weight.detach().numpy())
            remaining_para += np.sum(mask)
            layer.set_mask(mask)

    return remaining_para / total_para


# ---------- model ----------
class MaskableLinear(torch.nn.Linear):
    def __init__(self, in_features, out_features):
        self.mask = torch.ones(out_features, in_features)
        super().__init__(in_features, out_features)

    def set_mask(self, mask):
        self.mask = torch.tensor(mask)

    def forward(self, input):
        self.weight.data = self.weight.data * self.mask
        return super().forward(input)


model = torch.nn.Sequential(
    MaskableLinear(2, 100),
    torch.nn.ReLU(),
    MaskableLinear(100, 100),
    torch.nn.ReLU(),
    MaskableLinear(100, 1),
    torch.nn.Sigmoid(),
)

criterion = torch.nn.BCELoss()
optimizer = torch.optim.Adam(model.parameters(), weight_decay=0.001)

train()
# quantize()
accu = test()
# writer.add_scalars("pruning", {"accu": accu, "remaining": 1}, 0)

for i in range(20):
    remaining = pruning()
    train()
    # quantize()
    accu = test()
    print("remaining: %f, accu: %f" % (remaining, accu))
    # writer.add_scalars("pruning", {
    #     "accu": accu,
    #     "remaining": remaining
    # }, i + 1)
