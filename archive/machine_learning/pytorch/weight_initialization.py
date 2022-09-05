import matplotlib.pyplot as plt
import torch

N = 100
plt.style.use("default")

# plt.ylim(0, 40)


def init_weight(in_features, out_features):
    return torch.nn.Parameter(torch.randn(in_features, out_features))


class Layer(torch.nn.Module):
    def __init__(self, in_features, out_features, n):
        super().__init__()
        self.n = n
        self.w = init_weight(in_features, out_features)
        self.bn = torch.nn.BatchNorm1d(in_features)

    def forward(self, input):
        ret = torch.matmul(input, self.w)
        ret = torch.nn.functional.sigmoid(ret)
        ret = self.bn(ret)
        plt.subplot(1, 10, self.n)
        plt.hist(ret.detach().numpy().reshape(-1))
        return ret


def train():
    net = torch.nn.Sequential()

    for i in range(10):
        net.add_module("linear%d" % (i), Layer(N, N, i + 1))

    x = torch.rand(10, N)
    net(x)
    plt.show()


train()
