import numpy as np
import matplotlib.pyplot as plt
import torch

N = 1000
LAYERS = 30
plt.style.use("default")

ACTIVATION_FUNCTION = torch.nn.functional.sigmoid
WITH_BN = False


def init_weight(in_features, out_features):
    return torch.nn.Parameter(
        torch.randn(in_features, out_features) / np.sqrt(in_features)
    )


class Layer(torch.nn.Module):
    def __init__(self, in_features, out_features, n):
        super().__init__()
        self.n = n
        self.w = init_weight(in_features, out_features)
        self.bn = torch.nn.BatchNorm1d(N)

    def forward(self, input):
        ret = torch.matmul(input, self.w)
        if WITH_BN:
            ret = self.bn(ret)
        ret = ACTIVATION_FUNCTION(ret)
        self.output = ret
        return ret


def visualize():
    grads = [layer.w.grad.mean().item() for layer in net]
    grads = np.array(grads)
    grads *= 100
    plt.plot(grads, "bx")
    plt.show()


criterion = torch.nn.MSELoss()

net = torch.nn.Sequential()

for i in range(LAYERS):
    net.add_module("linear%d" % (i), Layer(N, N, i + 1))


def train():
    x = torch.rand(10, N)
    loss = criterion(net(x), torch.zeros(10, N))
    loss.backward()
    visualize()


train()
