function [J grad] = nnCostFunction(nn_params, ...
                                   input_layer_size, ...
                                   hidden_layer_size, ...
                                   num_labels, ...
                                   X, y, lambda)
%NNCOSTFUNCTION Implements the neural network cost function for a two layer
%neural network which performs classification
%   [J grad] = NNCOSTFUNCTON(nn_params, hidden_layer_size, num_labels, ...
%   X, y, lambda) computes the cost and gradient of the neural network. The
%   parameters for the neural network are "unrolled" into the vector
%   nn_params and need to be converted back into the weight matrices. 
% 
%   The returned parameter grad should be a "unrolled" vector of the
%   partial derivatives of the neural network.
%

% Reshape nn_params back into the parameters Theta1 and Theta2, the weight matrices
% for our 2 layer neural network
Theta1 = reshape(nn_params(1:hidden_layer_size * (input_layer_size + 1)), ...
                 hidden_layer_size, (input_layer_size + 1));

Theta2 = reshape(nn_params((1 + (hidden_layer_size * (input_layer_size + 1))):end), ...
                 num_labels, (hidden_layer_size + 1));

% Setup some useful variables
m = size(X, 1);
         
% You need to return the following variables correctly 
J = 0;
Theta1_grad = zeros(size(Theta1));
Theta2_grad = zeros(size(Theta2));

% ====================== YOUR CODE HERE ======================
% Instructions: You should complete the code by working through the
%               following parts.
%
% Part 1: Feedforward the neural network and return the cost in the
%         variable J. After implementing Part 1, you can verify that your
%         cost function computation is correct by verifying the cost
%         computed in ex4.m
%
% Part 2: Implement the backpropagation algorithm to compute the gradients
%         Theta1_grad and Theta2_grad. You should return the partial derivatives of
%         the cost function with respect to Theta1 and Theta2 in Theta1_grad and
%         Theta2_grad, respectively. After implementing Part 2, you can check
%         that your implementation is correct by running checkNNGradients
%
%         Note: The vector y passed into the function is a vector of labels
%               containing values from 1..K. You need to map this vector into a 
%               binary vector of 1's and 0's to be used with the neural network
%               cost function.
%
%         Hint: We recommend implementing backpropagation using a for-loop
%               over the training examples if you are implementing it for the 
%               first time.
%
% Part 3: Implement regularization with the cost function and gradients.
%
%         Hint: You can implement this around the code for
%               backpropagation. That is, you can compute the gradients for
%               the regularization separately and then add them to Theta1_grad
%               and Theta2_grad from Part 2.
%

## 1: costFunction

## Theta1: 25*401
## Theta2: 10*26
## X:      5000*401
## y:      5000*1

X=[ones(m,1) X];

for i=1:m
    a1=X(i,:);
    ty=y(i,:);
    z2=a1*Theta1';
    a2=sigmoid(z2);
    a2=[ones(1,1) a2];
    ## hypothesis: 1 *26
    z3=a2*Theta2';
    a3=sigmoid(z3);
    ## a2: 1 * 10
    delta3=zeros(num_labels,1);
    
    for k=1:num_labels
        delta3(k)=a3(k)-(ty==k);
        J = J-(ty==k)*log(a3(k))-(1-(ty==k))*log(1-a3(k));
    end
    delta2 = (delta3'*Theta2)(2:end) .* sigmoidGradient(z2);
    ## delta3: 10*1
    ## a2:     1*26
    ## 
    ## a1:     1*401
    ## delta2: 1*25
    Theta1_grad+=delta2'*a1;
    Theta2_grad+=delta3*a2;
end

Theta1_grad/=m;
Theta2_grad/=m;

temp=Theta1;
temp(:,1)=0;
Theta1_grad+=lambda*temp/m;

temp=Theta2;
temp(:,1)=0;
Theta2_grad+=lambda*temp/m;

J=J/m;

## 2: regularization

regulation=0;
for j=1:size(Theta1,1)
    for k=2:size(Theta1,2)
        regulation+=Theta1(j,k)^2;
    end
end

for j=1:size(Theta2,1)
    for k=2:size(Theta2,2)
        regulation+=Theta2(j,k)^2;
    end
end
regulation = regulation*lambda/2/m;

J+=regulation;

## 3: backpropagation

% -------------------------------------------------------------

% =========================================================================

% Unroll gradients
grad = [Theta1_grad(:) ; Theta2_grad(:)];


end
