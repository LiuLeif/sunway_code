package com.amazon.test;
import android.util.Log;
import android.test.InstrumentationTestCase;

// adb shell am instrument -w com.amazon.test/android.test.InstrumentationTestRunner
public class MainActivityTest extends InstrumentationTestCase {
    @Override
    public void setUp() {
        Log.d("sunway", "setup");
    }
    public void testCreateActivity() {
        assertEquals(1, 1);
    }
    public void testCreateActivity2() {
        assertEquals(2, 2);
    }
}
