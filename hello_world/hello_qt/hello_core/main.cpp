#include <QCoreApplication>
#include <QVector>

int main(int argc, char *argv[]) {
    QVector<int> data {1, 2, 3};
    // data.setSharable(false);
    printf("%p\n", &data[0]);

    auto dataCopy = data;
    printf("%p\n", &dataCopy[0]);

    return 0;
}
