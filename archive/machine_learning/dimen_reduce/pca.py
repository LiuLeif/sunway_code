import numpy as np
from numpy import linalg

X = np.array([[1, 2, 3, 4], [2, 4, 6, 8], [1, 3, 5, 7], [3, 9, 15, 21], [0, 1, 2, 3]])

# PCA with SVD

u, s, v = linalg.svd(X, full_matrices=False)

print (u.shape, s.shape, v.shape)
print(s)

X2 = np.dot(u * s, v)
print(X2)

u2 = u[:, :2]
Y = np.dot(u2.T, X)
print(Y)

# PCA with eigen decomposition
t = np.dot(X, X.T)
print(t)

s, v= linalg.eig(t)
print(s.shape, v.shape)
print(s)

t2 = np.dot(v*s,v.T)
print(t2)

v2 = v[:, :2]
Y = np.dot(v2.T, X)
print(Y)
