function [jVal, gradient] = costFunction(theta, X,y)
    jVal = computeCost(X,y,theta);
    m=2;
    gradient=zeros(m,1);
    for j=1:2,
        total=0;
        for i=1:m,
            total=total+((X(i,:)*theta)-y(i))*X(i,j);
        end
        gradient(j)=total/m;
    end
end
