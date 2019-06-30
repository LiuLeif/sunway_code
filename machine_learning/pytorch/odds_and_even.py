import numpy as np
from torch.utils.data import Dataset, DataLoader
import torch

N_CLASSES = 3
model = torch.nn.Sequential(
    torch.nn.Linear(1, 10), torch.nn.ReLU(), torch.nn.Linear(10, N_CLASSES)
)


class OddsAndEvenDataset(Dataset):
    def __init__(self, low, high, size):
        X = np.random.randint(low, high, size)
        Y = X % N_CLASSES
        self.X = torch.from_numpy(X).float().view(-1, 1)
        self.Y = torch.from_numpy(Y).long().view(-1)

    def __getitem__(self, index):
        return self.X[index], self.Y[index]

    def __len__(self):
        return len(self.X)


training_set = OddsAndEvenDataset(0, 1000, 500)
training_loader = DataLoader(training_set, batch_size=100)

test_set = OddsAndEvenDataset(500, 2000, 500)
test_loader = DataLoader(test_set, batch_size=500)

# criterion = torch.nn.BCEWithLogitsLoss()
criterion = torch.nn.CrossEntropyLoss()

optimizer = torch.optim.Adam(model.parameters(), lr=1e-3)


def train():
    model.train()
    for i in range(1000):
        for x, y in training_loader:
            loss = criterion(model(x), y)
            optimizer.zero_grad()
            loss.backward()
            optimizer.step()
        if i % 20 == 0:
            print("loss:", loss.item())


def test():
    model.eval()
    for x, y in training_loader:
        y_hat = model(x)
        # y_hat = F.sigmoid(y_hat)
        # y_hat = y_hat > 0.5
        y_hat = torch.argmax(y_hat, dim=1)
        accu = torch.sum(y_hat.byte() == y.byte()).item() / 100
        print("train:", accu)
        break

    for x, y in test_loader:
        y_hat = model(x)
        # y_hat = F.sigmoid(y_hat)
        # y_hat = y_hat > 0.5
        y_hat = torch.argmax(y_hat, dim=1)
        accu = torch.sum(y_hat.byte() == y.byte()).item() / 500
        print("test:", accu)


train()
test()
