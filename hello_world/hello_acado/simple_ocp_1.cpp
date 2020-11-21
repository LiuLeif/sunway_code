#include <acado_toolkit.hpp>
#include <acado_gnuplot.hpp>

int main() {
    USING_NAMESPACE_ACADO;

    DifferentialState x, v;
    Control            u;
    Parameter                T          ;
    DifferentialEquation     f(0.0, T);

    OCP ocp(0.0, T);
    ocp.minimizeMayerTerm(T);
    // ocp.minimizeLagrangeTerm(T);

    f << dot(x) == v;
    f << dot(v) == u;

    ocp.subjectTo(f);
    ocp.subjectTo(AT_START, x ==  0.0);
    ocp.subjectTo(AT_START, v ==  0.0);
    ocp.subjectTo(AT_START, u ==  0.0);

    ocp.subjectTo(AT_END, x == 20.0);
    ocp.subjectTo(AT_END, v ==  0.0);

    ocp.subjectTo(-1 <= u <= 1);
    ocp.subjectTo(0 <= T <= 100.0);

    GnuplotWindow window;
    window.addSubplot(x, "THE DISTANCE s");
    window.addSubplot(v, "THE VELOCITY v");
    window.addSubplot(u, "THE CONTROL INPUT u");

    OptimizationAlgorithm algorithm(ocp);
    algorithm << window;
    algorithm.solve();

    return 0;
}
