#include "message_box.h"

#include <QDebug>

MessageBox::MessageBox(QObject *parent) : QObject(parent) {
}

void MessageBox::onTextChanged(QString text) {
    qDebug() << "onTextChanged " << text;
}
