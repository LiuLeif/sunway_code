package com.amazon.test;

import android.app.Activity;
import android.content.Intent;
import android.util.Log;
import android.os.Bundle;
import android.app.Instrumentation;

// adb shell am instrument -w com.amazon.test/.MyInstrumentation
public class MyInstrumentation extends Instrumentation {
    @Override
    public void onCreate(Bundle args) {
        super.onCreate(args);
        Log.d("sunway", "onCreate");
        start();
    }

    @Override
    public void onStart() {
        Log.d("sunway", "onStart");
        // 2. start activity
        Intent intent = new Intent(getTargetContext(), MainActivity.class);
        intent.addFlags(Intent.FLAG_ACTIVITY_NEW_TASK | Intent.FLAG_ACTIVITY_CLEAR_TASK);
        startActivitySync(intent);
    }

    @Override
    public void callActivityOnCreate(Activity activity, Bundle icicle) {
        // 1. hook Activity.onCreate
        super.callActivityOnCreate(activity, icicle);
        Log.d("sunway", "callActivityOnCreate");
    }
}
