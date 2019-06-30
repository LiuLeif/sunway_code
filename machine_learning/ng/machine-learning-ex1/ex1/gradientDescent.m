function [theta, J_history] = gradientDescent(X, y, theta, alpha, num_iters)
%GRADIENTDESCENT Performs gradient descent to learn theta
%   theta = GRADIENTDESCENT(X, y, theta, alpha, num_iters) updates theta by 
%   taking num_iters gradient steps with learning rate alpha

% Initialize some useful values
    m = length(y); % number of training examples
    J_history = zeros(num_iters, 1);

    options=optimset('MaxIter','1000')
    [optTheta, val,exitFlag]=fminunc(@(t)(costFunction(t,X,y)),theta,options)
    theta=optTheta
end

function [theta, J_history] = gradientDescent2(X, y, theta, alpha, num_iters)
    %GRADIENTDESCENT Performs gradient descent to learn theta
    %   theta = GRADIENTDESCENT(X, y, theta, alpha, num_iters) updates theta by 
    %   taking num_iters gradient steps with learning rate alpha

                                % Initialize some useful values
    m = length(y); % number of training examples
    J_history = zeros(num_iters, 1);

    for iter = 1:num_iters

          % ====================== YOUR CODE HERE ======================
          % Instructions: Perform a single gradient step on the parameter vector
          %               theta. 
          %
          % Hint: While debugging, it can be useful to print out the values
          %       of the cost function (computeCost) and gradient here.
          %
        tmp=theta;

        for j=1:2,
            total=0;
            for i=1:m,
                ## total=total+(theta(1)*x(i,1)+theta(2)*x(i,2)-y(i))*x(i,j);
                total=total+((X(i,:)*theta)-y(i))*X(i,j);
            end
            tmp(j)=tmp(j)-alpha*total/m;
        end
        
        theta=tmp;
                  % ============================================================

                                % Save the cost J in every iteration    
        J_history(iter) = computeCost(X, y, theta);

    end

end
