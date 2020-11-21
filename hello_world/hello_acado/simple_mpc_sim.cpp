#include <acado_toolkit.hpp>
#include <acado_gnuplot.hpp>


int main() {
    USING_NAMESPACE_ACADO;

    //  _____
    // < OCP >
    //  -----
    //         \   ^__^
    //          \  (oo)\_______
    //             (__)\       )\/\
    //                 ||----w |
    //                 ||     ||
    DifferentialState v;
    Control u;

    DifferentialEquation f;
    f << dot(v) == u;

    OCP ocp(0.0, 1.0, 10);

    Function h;
    h << v;

    DVector r(1);
    r(0) = 10;

    ocp.minimizeLSQ(h, r);
    ocp.subjectTo(f);
    ocp.subjectTo(-20 <= u <= 20);

    //  ____________
    // < controller >
    //  ------------
    //         \   ^__^
    //          \  (oo)\_______
    //             (__)\       )\/\
    //                 ||----w |
    //                 ||     ||
    RealTimeAlgorithm alg(ocp, 0.1);

    StaticReferenceTrajectory reference;
    Controller controller(alg, reference);

    //  _____________
    // < environment >
    //  -------------
    //         \   ^__^
    //          \  (oo)\_______
    //             (__)\       )\/\
    //                 ||----w |
    //                 ||     ||
    OutputFcn g;
    Disturbance R;
    g << v * (1 + R);

    DynamicSystem dynamicSystem(f, g);
    Process process(dynamicSystem, INT_RK45);
    VariablesGrid disturbance;
    disturbance.read("mpc_dist.txt");
    if (process.setProcessDisturbance(disturbance) != SUCCESSFUL_RETURN)
        exit(EXIT_FAILURE);
    SimulationEnvironment sim(0.0, 5.0, process, controller);
    DVector x0(1);
    x0(0) = 0;

    if (sim.init(x0) != SUCCESSFUL_RETURN)
        exit(EXIT_FAILURE);
    if (sim.run() != SUCCESSFUL_RETURN)
        exit(EXIT_FAILURE);

    VariablesGrid diffStates;
    sim.getProcessDifferentialStates(diffStates);

    VariablesGrid feedbackControl;
    sim.getFeedbackControl(feedbackControl);

    GnuplotWindow window;
    window.addSubplot(diffStates(0),   "v");
    window.addSubplot(feedbackControl, "u");

    window.plot();

    return 0;
}
