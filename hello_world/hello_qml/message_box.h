#ifndef MESSAGE_H
#define MESSAGE_H


#include <QObject>

class MessageBox : public QObject {
    Q_OBJECT

  public:
    explicit MessageBox(QObject *parent = 0);

  public slots:
    void onTextChanged(QString);
};

#endif  // MESSAGE_H
