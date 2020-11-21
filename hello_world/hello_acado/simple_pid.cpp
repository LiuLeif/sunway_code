#include <acado_toolkit.hpp>
#include <acado_gnuplot.hpp>

int main( ) {
    USING_NAMESPACE_ACADO

    PIDcontroller pid( 4, 1, 0.01 );

    DVector pWeights( 4 );
    pWeights(0) = 1.0;
    pWeights(1) = -1000.0;
    pWeights(2) = 1000.0;
    pWeights(3) = -1000.0;

    DVector iWeights( 4 );
    iWeights(0) = 2.0;
    iWeights(1) = 0.0;
    iWeights(2) = 20.0;
    iWeights(3) = -20.0;

    pid.setProportionalWeights( pWeights );
    pid.setIntegralWeights( iWeights );

    pid.setControlLowerLimit( 0, -200.0 );
    pid.setControlUpperLimit( 0, 200.0 );

    StaticReferenceTrajectory zeroReference("pid_ref.txt");

    Controller controller( pid, zeroReference );

    DVector y( 4 );
    y.setZero( );
    y(0) = 0;

    controller.init( 0.0, y );
    controller.step( 0.0, y );

    DVector u;
    controller.getU( u );
    printf("%f\n", u(0));
    return 0;
}
