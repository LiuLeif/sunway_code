// 2019-07-27 20:40
import static org.mockito.Mockito.*;
import java.util.List;
import java.util.ArrayList;
import org.junit.Test;
import static org.junit.Assert.*;

public class TestSample {
    @Test
    public void testMock() {
        List<Integer> x = mock(ArrayList.class);
        when (x.size()).thenReturn(1);
        assertEquals(x.size(), 1);
        x.add(1);
        assertEquals(x.get(0), null);
    }

    @Test
    public void testSpy() {
        List<Integer> x = spy(ArrayList.class);
        x.add(1);
        doReturn(2).when(x).size();
        assertEquals(x.size(), 2);
        assertEquals(x.get(0).intValue(), 1);
    }

    @Test
    public void testAnnotation() {
        @Mock
        ArrayList<Integer> x;
        when (x.size()).thenReturn(1);
        assertEquals(x.size(), 1);
    }
}
