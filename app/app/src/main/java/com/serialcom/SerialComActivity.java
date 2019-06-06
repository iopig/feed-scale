package com.serialcom;

import android.graphics.Color;
import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.DividerItemDecoration;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.view.View;
import android.widget.AdapterView;
import android.widget.Button;
import android.widget.EditText;
import android.widget.RadioGroup;
import android.widget.Spinner;
import android.widget.TextView;
import android.widget.Toast;

/**********study UART **********/
import com.bjw.bean.ComBean;
import com.bjw.utils.FuncUtil;
import com.bjw.utils.SerialHelper;

import java.io.IOException;

import android_serialport_api.SerialPortFinder;

public class SerialComActivity extends AppCompatActivity {

    private RecyclerView receDisp;
    private Spinner spUARTList;

    private EditText edInputdata;
    private Button btSenddata;
    private Button btZero;
    private Button btMeas;
    private Button btUploading;
    private Button btExit;

    private Button btCottage;

    private TextView tvFarmer;
    private TextView tvPig;
    private TextView tvFodder;
    private TextView tvSuggestion;

    private RadioGroup radioGroup;

    private SerialPortFinder serialPortFinder;
    private SerialHelper serialHelper;

    private Spinner spBitRate;
    private Button btOpenPort;

    private LogListAdapter logListAdapter;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_seria_com);

        receDisp = (RecyclerView) findViewById(R.id.receDisp);
        //spUARTList = (Spinner) findViewById(R.id.sp_serial);
        edInputdata = (EditText) findViewById(R.id.ed_input);
        btSenddata = (Button) findViewById(R.id.bt_send);
        btOpenPort = (Button) findViewById(R.id.bt_open);
        btZero = (Button) findViewById(R.id.zero);
        btMeas  = (Button)findViewById(R.id.measure);
        btUploading = (Button)findViewById(R.id.uploading);
        btExit  = (Button)findViewById(R.id.exit);

        tvFarmer    = (TextView)findViewById(R.id.farmer);
        tvFarmer.setText("家庭农场:陈东");
        tvFarmer.setTextSize(24);
        tvFarmer.setBackgroundColor(Color.parseColor("#00FF00"));

        tvPig       = (TextView)findViewById(R.id.pig);
        tvPig.setText("外三PIC01");
        tvPig.setTextSize(24);
        tvPig.setBackgroundColor(Color.parseColor("#F0F000"));

        tvFodder    = (TextView)findViewById(R.id.fodder);
        tvFodder.setText("永祥021");
        tvFodder.setTextSize(24);
        tvFodder.setBackgroundColor(Color.parseColor("#40F000"));

        tvSuggestion= (TextView)findViewById(R.id.suggestion);
        tvSuggestion.setText("20kg");
        tvSuggestion.setTextSize(24);
        tvSuggestion.setBackgroundColor(Color.parseColor("#F01010"));

        radioGroup = (RadioGroup) findViewById(R.id.radioGroup);


        logListAdapter = new LogListAdapter(null);

        receDisp.setLayoutManager(new LinearLayoutManager(this));
        receDisp.setAdapter(logListAdapter);
        receDisp.addItemDecoration(new DividerItemDecoration(this, DividerItemDecoration.VERTICAL));

        /*Here we set an array to set all the job CARDs*/
        final Integer[]cotId = new Integer[]{R.id.left1,R.id.left2,R.id.left3,R.id.left4,R.id.left5,
                                       R.id.left6,R.id.left7,R.id.left8,R.id.left9,R.id.left10,
                                    R.id.right1,R.id.right2,R.id.right3,R.id.right4,R.id.right5,
                                    R.id.right6,R.id.right7,R.id.right8,R.id.right9,R.id.right10};
   //     Button bt101,bt102,bt103,bt104,bt105,bt106,bt107,bt108,bt109,bt110;
    //    Button bt201,bt202,bt203,bt204,bt205,bt206,bt207,bt208,bt209,bt210;
//        Button[] btCots = new Button[]{bt101,bt102,bt103,bt104,bt105,bt106,bt107,bt108,bt109};
        for (int nuCot = 0; nuCot < 20; nuCot++){
            btCottage = (Button)findViewById(cotId[nuCot]);
            //btCottage.setText("Good Job!");
            //btCottage.setBackgroundColor(Color.parseColor("#A0A0A0"));

            btCottage.setOnClickListener(new View.OnClickListener(){

                    public void onClick(View v){
 //                       Button btCotId;
                        //btCotId = (Button)findViewById(cotId[v.getId()]);
                        //v.setText("Done!");
                        v.setBackgroundColor(Color.parseColor("#0E7012"));
                    }
            });
        }
        iniview();
    }


    @Override
    protected void onDestroy() {
        super.onDestroy();
        serialHelper.close();
    }

    private void iniview() {


        serialPortFinder = new SerialPortFinder();
        serialHelper = new SerialHelper() {
            @Override
            protected void onDataReceived(final ComBean comBean) {
                runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        Toast.makeText(getBaseContext(), FuncUtil.ByteArrToHex(comBean.bRec), Toast.LENGTH_SHORT).show();
                        logListAdapter.addData(comBean.sRecTime + ":   " + (new String(comBean.bRec)));//Print the sting intstead of HEX
                        receDisp.smoothScrollToPosition(logListAdapter.getData().size());
                    }
                });
            }
        };

        final String[] ports = serialPortFinder.getAllDevicesPath();
        final String[] bitrates = new String[]{"9600", "19200", "38400", "57600", "115200", "230400", "460800", "500000", "576000", "921600", "1000000", "1152000", "1500000", "2000000", "2500000", "3000000", "3500000", "4000000"};



        //SpAdapter spAdapter = new SpAdapter(this);
        //spAdapter.setDatas(ports);
        //spUARTList.setAdapter(spAdapter);
/*
        spUARTList.setOnItemSelectedListener(new AdapterView.OnItemSelectedListener() {
            @Override
            public void onItemSelected(AdapterView<?> parent, View view, int position, long id) {
                serialHelper.close();
                serialHelper.setPort(ports[position]);
                btOpenPort.setEnabled(true);
            }

            @Override
            public void onNothingSelected(AdapterView<?> parent) {

            }
        });
*/

        /**********Open the UART *************/
        btOpenPort.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                try {
                    serialHelper.close();
                    if (radioGroup.getCheckedRadioButtonId() == R.id.radioButton1){ //Sent text
                        serialHelper.setPort("/dev/ttyS1");
                        Toast.makeText(getBaseContext(), "Select the com[1]", Toast.LENGTH_SHORT).show();
                    }else{
                        serialHelper.setPort("/dev/ttyS4");
                        Toast.makeText(getBaseContext(), "Select the com[4]", Toast.LENGTH_SHORT).show();
                    }

                   // btOpenPort.setEnabled(true);
                    serialHelper.open();
                    Toast.makeText(getBaseContext(), "Open uart once you change the com selection", Toast.LENGTH_SHORT).show();
                    //btOpenPort.setEnabled(false);
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        });


        /**********Send command to the electronic scale*************/
        btSenddata.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {

                if (edInputdata.getText().toString().length() > 0) {
                    if (serialHelper.isOpen()) {
                        serialHelper.sendTxt(edInputdata.getText().toString());
                    } else {
                        Toast.makeText(getBaseContext(), "Please, open the com port", Toast.LENGTH_SHORT).show();
                    }
                } else {
                    Toast.makeText(getBaseContext(), "Fill the data to send", Toast.LENGTH_SHORT).show();
                }
//                  } else {//hex format
            }
        });

        /**********adjust the electronic scale to Zero*************/
        btZero.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if (serialHelper.isOpen()) {
                    serialHelper.sendTxt("T");
                } else {
                    Toast.makeText(getBaseContext(), "Please, open the com port", Toast.LENGTH_SHORT).show();
                }
            }
        });

        btMeas.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if (serialHelper.isOpen()) {
                    serialHelper.sendTxt("R");
                } else {
                    Toast.makeText(getBaseContext(), "Please, open the com port", Toast.LENGTH_SHORT).show();
                }
            }
        });
    }
}
