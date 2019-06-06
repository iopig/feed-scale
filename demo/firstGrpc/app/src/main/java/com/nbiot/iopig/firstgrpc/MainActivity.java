package com.nbiot.iopig.firstgrpc;

import android.os.Bundle;
import android.support.design.widget.FloatingActionButton;
import android.support.design.widget.Snackbar;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.Toolbar;
import android.view.View;
import android.view.Menu;
import android.view.MenuItem;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import  cn.iopig.feed.scale.grpc.*;



public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        Toolbar toolbar = (Toolbar) findViewById(R.id.toolbar);
        setSupportActionBar(toolbar);
        grpctest();
        FloatingActionButton fab = (FloatingActionButton) findViewById(R.id.fab);
        fab.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                Snackbar.make(view, "Replace with your own action", Snackbar.LENGTH_LONG)
                        .setAction("Action", null).show();
            }
        });
    }


    protected String grpctest() {
        String host = "117.139.13.149";
        String portStr ="50051";
        int port =  Integer.valueOf(portStr);
        try {
            ManagedChannel channel = ManagedChannelBuilder.forAddress(host, port).usePlaintext().build();

            FsPadGrpc.FsPadBlockingStub stub = FsPadGrpc.newBlockingStub(channel);

            ReqHeader reqHeader = ReqHeader.newBuilder().setVersion(1).setDevId("1").build();
            DevInfoReq request = DevInfoReq.newBuilder().setReqHeader(reqHeader).build();
            PigstyInfoRes reply = stub.padLogin(request);
            return  reply.getFarmer();
            //return reply.getMessage();

          // return "sdfsdf";
        } catch (Exception e) {
            return "Failed... ";
        }
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.menu_main, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        // Handle action bar item clicks here. The action bar will
        // automatically handle clicks on the Home/Up button, so long
        // as you specify a parent activity in AndroidManifest.xml.
        int id = item.getItemId();

        //noinspection SimplifiableIfStatement
        if (id == R.id.action_settings) {
            return true;
        }

        return super.onOptionsItemSelected(item);
    }
}
