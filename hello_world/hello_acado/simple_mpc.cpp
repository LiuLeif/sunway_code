#include <acado_toolkit.hpp>
#include <acado_gnuplot.hpp>

int main( ) {
    USING_NAMESPACE_ACADO

    DifferentialState x, v;
    Control u;

    DifferentialEquation f;
    f << dot(x) == v;
    f << dot(v) == u;

    OCP ocp(0, 1, 20);
    ocp.minimizeMayerTerm((v - 20) * (v - 20));
    ocp.subjectTo(f);
    ocp.subjectTo(-1 <= u <= 1);


    OCPexport mpc( ocp );
    mpc.set( HESSIAN_APPROXIMATION,       EXACT_HESSIAN  		);
    mpc.set( DISCRETIZATION_TYPE,         MULTIPLE_SHOOTING 	);
    mpc.set( INTEGRATOR_TYPE,             INT_RK4   			);
    mpc.set( NUM_INTEGRATOR_STEPS,        18            		);
    mpc.set( QP_SOLVER,                   QP_QPOASES    		);
    mpc.set( HOTSTART_QP,                 NO             		);
    mpc.set( GENERATE_TEST_FILE,          YES            		);
    mpc.set( GENERATE_MAKE_FILE,          YES            		);
    mpc.set( GENERATE_MATLAB_INTERFACE,   NO            		);
    mpc.set( SPARSE_QP_SOLUTION, 		  FULL_CONDENSING_N2	);
    mpc.set( DYNAMIC_SENSITIVITY, 		  SYMMETRIC				);
    mpc.set( CG_HARDCODE_CONSTRAINT_VALUES, NO 					);
    mpc.set( CG_USE_VARIABLE_WEIGHTING_MATRIX, YES 				);
    mpc.exportCode("mpc_export");
    mpc.printDimensionsQP( );

    return 0;
}
