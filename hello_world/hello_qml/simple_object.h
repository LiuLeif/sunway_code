// 2019-12-31 18:11
#ifndef SIMPLE_OBJECT_H
#define SIMPLE_OBJECT_H

#include <QObject>

class SimpleObject : public QObject {
    Q_OBJECT

  public:
    explicit SimpleObject(QObject *parent = 0);

  signals:


  public slots:

};


#endif  // SIMPLE_OBJECT_H
