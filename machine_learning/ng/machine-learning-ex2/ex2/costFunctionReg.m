function [J, grad] = costFunctionReg(theta, X, y, lambda)
%COSTFUNCTIONREG Compute cost and gradient for logistic regression with regularization
%   J = COSTFUNCTIONREG(theta, X, y, lambda) computes the cost of using
%   theta as the parameter for regularized logistic regression and the
%   gradient of the cost w.r.t. to the parameters. 

% Initialize some useful values
m = length(y); % number of training examples

% You need to return the following variables correctly 
J = 0;
grad = zeros(size(theta));

% ====================== YOUR CODE HERE ======================
% Instructions: Compute the cost of a particular choice of theta.
%               You should set J to the cost.
%               Compute the partial derivatives and set grad to the partial
%               derivatives of the cost w.r.t. each parameter in theta


function h = hypothesis(theta,x)
    h=sigmoid(x*theta);
end

total=0;
for i=1:m
    total+=-y(i)*log(hypothesis(theta,X(i,:)))-(1-y(i))*log(1-hypothesis(theta,X(i,:)));
endfor

J=total/m+lambda*sum(theta .^ 2)/(2*m)-lambda*(theta(1)^2)/(2*m);

total=0;
for i=1:m
    total+=(hypothesis(theta,X(i,:))-y(i))*X(i,1);
endfor
grad(1)=total/m;

for j=2:size(grad)
    total=0;
    for i=1:m
        total+=(hypothesis(theta,X(i,:))-y(i))*X(i,j);
    endfor
    grad(j)=total/m+lambda*theta(j)/m;
endfor    

% =============================================================

end
