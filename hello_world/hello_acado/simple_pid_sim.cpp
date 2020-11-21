#include <acado_toolkit.hpp>
#include <acado_gnuplot.hpp>


int main() {
    USING_NAMESPACE_ACADO;

    //  _______
    // < model >
    //  -------
    //         \   ^__^
    //          \  (oo)\_______
    //             (__)\       )\/\
    //                 ||----w |
    //                 ||     ||
    DifferentialState v;
    Control u;

    DifferentialEquation f;
    f << dot(v) == u;

    //  ____________
    // < controller >
    //  ------------
    //         \   ^__^
    //          \  (oo)\_______
    //             (__)\       )\/\
    //                 ||----w |
    //                 ||     ||
    // controller
    PIDcontroller pid(1, 1, 0.1);

    DVector pWeights(1);
    pWeights(0) = 5;

    DVector iWeights(1);
    iWeights(0) = 5.0;

    DVector dWeights(1);
    dWeights(0) = 0;

    pid.setProportionalWeights(pWeights);
    pid.setIntegralWeights(iWeights);
    pid.setDerivativeWeights(dWeights);

    pid.setControlLowerLimit(0, -20);
    pid.setControlUpperLimit(0, 20);

    VariablesGrid data(1, 0.0, 1.0, 2);
    data(0, 0) = 5;
    data(1, 0) = 5;

    PeriodicReferenceTrajectory reference(data);

    Controller controller(pid, reference);

    //  _____________
    // < environment >
    //  -------------
    //         \   ^__^
    //          \  (oo)\_______
    //             (__)\       )\/\
    //                 ||----w |
    //                 ||     ||
    Disturbance R;
    OutputFcn g;
    g << v * (1 + R);
    DynamicSystem dynamicSystem(f, g);

    Process process(dynamicSystem, INT_RK45);

    VariablesGrid disturbance;
    disturbance.read("pid_dist.txt");
    if (process.setProcessDisturbance(disturbance) != SUCCESSFUL_RETURN)
        exit(EXIT_FAILURE);

    SimulationEnvironment sim(0.0, 5.0, process, controller);

    DVector x0(1);
    x0.setZero();

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
    window.addSubplot(disturbance,     "disturbance");
    window.plot();

    return 0;
}
