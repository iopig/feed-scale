package com.serialcom;

import android.app.Activity;
import android.app.admin.DevicePolicyManager;
import android.content.ComponentName;
import android.content.Context;
import android.content.Intent;
import android.graphics.Color;
import android.icu.text.SymbolTable;
import android.os.Bundle;
import android.os.Environment;
import android.os.Handler;
import android.os.Message;
import android.provider.Settings;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.DividerItemDecoration;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.text.TextUtils;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.RadioGroup;
import android.widget.TextView;
import android.widget.Toast;

/**********UART Part**********/
import com.bjw.bean.ComBean;
import com.bjw.utils.SerialHelper;
import com.google.protobuf.InvalidProtocolBufferException;
import com.serialcom.FarmManager.buildingInfo;

import java.io.BufferedWriter;
import java.io.ByteArrayInputStream;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;

import android_serialport_api.SerialPortFinder;
/************gRPC Part **********/
import cn.iopig.feed.scale.grpc.DevInfoReq;
import cn.iopig.feed.scale.grpc.FsPadGrpc;
import cn.iopig.feed.scale.grpc.PigHouseInfo;
import cn.iopig.feed.scale.grpc.PigstyInfo;
import cn.iopig.feed.scale.grpc.PigstyInfoRes;
import cn.iopig.feed.scale.grpc.ReqHeader;
import cn.iopig.feed.scale.grpc.ScaleDevRawData;
import cn.iopig.feed.scale.grpc.UploadDevDateReq;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.opencensus.internal.Utils;

/************Timer Part **********/
import java.io.InputStream;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;
import java.io.OutputStreamWriter;
import java.lang.ref.WeakReference;
import java.text.DateFormat;
import java.text.DecimalFormat;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Calendar;
import java.util.Date;
import java.util.List;
import java.util.Locale;
import java.util.TimeZone;
import java.util.Timer;
import java.util.TimerTask;
import java.util.UUID;

import static cn.iopig.feed.scale.grpc.FeedType.CHOOSE_PIGSTY;
import static cn.iopig.feed.scale.grpc.FeedType.LOAD;

class ColumnData{
    int colID;
    double suggWeight;
    double finishedWeight;
    double WeightMark; //record the weight water mark on the column selected time

    public ColumnData(int i, double v, double v1, double v2) {
        colID = i;
        suggWeight = v;
        finishedWeight = v1;
        WeightMark = v2;
    }
};


public class SerialComActivity extends AppCompatActivity {

    private WeightBuffer weightDataBuf;
    private FarmManager farmRunConfig;
    private ColumnData[] coData;
    private int activeColumn;
    private boolean isAddingMode = true;
    private int countDebug = 1;
    private boolean isRequestServerConfig;// one time/day

    private RecyclerView receDisp;

    private EditText edInputdata;
    private Button btSenddata;
    private Button btZero;
    private Button btMeas;
    private Button btUploading;//TODO
    private Button btExit;//TODO

    private Button btCottage;

    private TextView tvFarmer;
    private TextView tvPig;
    private TextView tvFodder;
    private TextView tvSuggestion;
    private TextView tvActiveCol;

    private RadioGroup radioGroup;// debug only

    private SerialHelper serialScaler;
    private SerialHelper serialNBiot;

    private Button btOpenPort;// debug only

    private LogListAdapter logListAdapter;

    /************Here we set an array to set all task cards*/
    public final Integer[]cotId;

    public SerialComActivity() {
        weightDataBuf = new WeightBuffer(1024);
        //default is adding mode
        cotId = new Integer[]{
                R.id.left1,R.id.right1,
                R.id.left2,R.id.right2,
                R.id.left3,R.id.right3,
                R.id.left4,R.id.right4,
                R.id.left5,R.id.right5,
                R.id.left6,R.id.right6,
                R.id.left7,R.id.right7,
                R.id.left8,R.id.right8,
                R.id.left9,R.id.right9,
                R.id.left10,R.id.right10};
        //TODO **debug only start, should get from the server
        coData = new ColumnData[20];
        coData[0] = new ColumnData(1,28.1,0.0,0.1);
        coData[1] = new ColumnData(2,2.2,2.0,0.2);
        coData[2] = new ColumnData(3,2.3,2.0,0.3);
        coData[3] = new ColumnData(4,3.4,1.0,2.4);
        coData[4] = new ColumnData(5,4.5,2.0,2.5);
        coData[5] = new ColumnData(6,5.6,2.0,2.6);
        coData[6] = new ColumnData(7,2.7,2.0,2.7);
        coData[7] = new ColumnData(8,2.8,2.0,2.8);
        coData[8] = new ColumnData(9,2.9,2.0,2.9);
        coData[9] = new ColumnData(10,3.0,3.0,3.0);
        coData[10] = new ColumnData(11,3.0,3.0,3.0);
        coData[11] = new ColumnData(12,3.0,3.0,3.0);
        coData[12] = new ColumnData(13,3.1,3.0,3.9);
        coData[13] = new ColumnData(14,3.2,3.0,3.8);
        coData[14] = new ColumnData(15,3.3,3.0,3.4);
        coData[15] = new ColumnData(16,3.4,3.0,3.2);
        coData[16] = new ColumnData(17,3.5,3.0,3.2);
        coData[17] = new ColumnData(18,3.6,3.0,3.3);
        coData[18] = new ColumnData(19,3.7,3.0,3.2);
        coData[19] = new ColumnData(20,3.8,3.0,3.1);
        //**debug end
        isRequestServerConfig = true;
        farmRunConfig = new FarmManager();
    }

    private void SaveConfigtoFile(String filename, FarmManager configData) {

        FileOutputStream out = null;
        File dir;
        dir = getFilesDir();
        File dst;
        try {
            dst = new File(dir, filename);
            out = new FileOutputStream(dst);
            ObjectOutputStream oStream = new ObjectOutputStream(out);
            oStream.writeObject(configData);
            oStream.flush();
            oStream.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private void SaveByestoFile(String filename, byte[] buf, int offset, int size) {

        FileOutputStream out = null;
        File dir;
        dir = getFilesDir();
        File dst;
        try {
            dst = new File(dir, filename);
            out = new FileOutputStream(dst);
            out.write(buf,offset, size);
            out.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
    private byte[] LoadBytesfromFile(String filename){
        File dir;
        dir = getFilesDir();
        File dst;
       try{
           dst = new File(dir, filename);
           InputStream inStream = new FileInputStream(dst);
           int size=inStream.available();
           byte[] buf = new byte[size] ;
           inStream.read(buf,0,size);
           inStream.close();
           return buf;
       }catch (IOException e){
           e.printStackTrace();
       }
       return null;
    }

    private boolean CheckFileExist(String filename){
        File dir;
        dir = getFilesDir();
        File dst;
        dst = new File(dir, filename);
        if (dst.exists())
            return true;
        else
            return false;
    }

    private static void ListDirectory(String dirname) // For debug only
    {
        File f1 = new File(dirname);
        if (f1.isDirectory()) {
            System.out.println("IOPIG debug Directory " + dirname);
            String s[] = f1.list();
            for (int i = 0; i < s.length; i++) {
                File f = new File(dirname + "/" + s[i]);
                if (f.isDirectory()) {
                    System.out.println(s[i] + " is a directory");
                } else {
                    System.out.println(s[i] + " is a file");
                }
            }
        } else {
            System.out.println(dirname + " not a directory");
        }
    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_seria_com);

        /*********for debug only *************/
        receDisp = (RecyclerView) findViewById(R.id.receDisp);
        edInputdata = (EditText) findViewById(R.id.ed_input);
        edInputdata.setVisibility(edInputdata.INVISIBLE);
        btSenddata = (Button) findViewById(R.id.bt_send);
        btSenddata.setVisibility(btSenddata.INVISIBLE);
        btOpenPort = (Button) findViewById(R.id.bt_open);
        btOpenPort.setVisibility(btOpenPort.INVISIBLE);

        /*********basic control *************/
        btZero = (Button) findViewById(R.id.zero);
        btMeas  = (Button)findViewById(R.id.measure);
        btMeas.setText("");
        btUploading = (Button)findViewById(R.id.uploading);
        btExit  = (Button)findViewById(R.id.exit);
        btExit.setText("");

        /*********display the information for the farmer****/
        tvFarmer    = (TextView)findViewById(R.id.farmer);
        tvFarmer.setText("Device is OFFLINE");
        tvFarmer.setGravity(View.TEXT_ALIGNMENT_GRAVITY);
        tvFarmer.setTextSize(36);
        tvFarmer.setBackgroundColor(Color.parseColor("#A0FA40"));

        tvPig       = (TextView)findViewById(R.id.pig);
        tvPig.setText("日龄XXX天");
        tvPig.setTextSize(32);
        tvPig.setGravity(View.TEXT_ALIGNMENT_GRAVITY);
        tvPig.setBackgroundColor(Color.parseColor("#F0F000"));

        tvFodder    = (TextView)findViewById(R.id.fodder);
        tvFodder.setText("饲料规格");
        tvFodder.setGravity(View.TEXT_ALIGNMENT_GRAVITY);
        tvFodder.setTextSize(32);
        tvFodder.setBackgroundColor(Color.parseColor("#40F000"));

        tvSuggestion= (TextView)findViewById(R.id.suggestion);
        tvSuggestion.setText("上料状态");
        tvSuggestion.setTextSize(36);
        tvSuggestion.setGravity(View.TEXT_ALIGNMENT_GRAVITY);
        tvSuggestion.setBackgroundColor(Color.parseColor("#A0FA40"));

        tvActiveCol = (TextView)findViewById(R.id.activeColumn);
        tvActiveCol.setText("请选择猪栏");
        tvActiveCol.setGravity(View.TEXT_ALIGNMENT_GRAVITY);
        tvActiveCol.setTextSize(36);
        tvActiveCol.setBackgroundColor(Color.parseColor("#F0AA70"));


        /*********for debug only *************/
        radioGroup = (RadioGroup) findViewById(R.id.radioGroup);
        radioGroup.setVisibility(radioGroup.INVISIBLE);
        logListAdapter = new LogListAdapter(null);
        receDisp.setLayoutManager(new LinearLayoutManager(this));
        receDisp.setAdapter(logListAdapter);
        receDisp.addItemDecoration(new DividerItemDecoration(this, DividerItemDecoration.VERTICAL));

        initClickFunc();

        //Send command, by the serial port[5G NBiot]
        initNBiot();
        //grpctest();// Communicate the server, by 4G/WIFI

        initWeightScaler();
        initTimer();


        //Load the latest config from local file,  or  wait from the server,in NBiot mode it may takes much more seconds
        /*While(1){}
         * Read the file from filesystem to farmRunConfig
         * Validate the file
         * if correct,break;
         */

        if (CheckFileExist("AFS20190616.cfg")) {
            byte[] seriDataBuf;
            seriDataBuf = LoadBytesfromFile("AFS20190616.cfg");
            if ((seriDataBuf == null)) throw new AssertionError();

            PigstyInfoRes newFarmConfig= null;
            try {
                newFarmConfig = PigstyInfoRes.parseFrom(seriDataBuf);
            } catch (InvalidProtocolBufferException e) {
                e.printStackTrace();
            }
            if(newFarmConfig != null){
                if(farmRunConfig.updateFarmInfo(newFarmConfig)){
                    tvFarmer.setText("Config file is OK");
                    // Update from local file
                    initCottage();
                }
            }

            if(farmRunConfig != null){
                isRequestServerConfig = false;//no need to requset the server
            }
            else{
                isRequestServerConfig = true;
                tvFarmer.setText("首次初始化，等待完成后重启");
            }
        }else{
                //ListDirectory("/data/data/com.serialcom/files");
            isRequestServerConfig = true;
            tvFarmer.setText("首次初始化，等待完成后重启");
         }
    }

    private void initCottage(){
        tvFarmer.setText(farmRunConfig.getFarmer());

        // switch the building, by button input?

        buildingInfo bInfo = new buildingInfo();
        bInfo = farmRunConfig.buildingArrayInfo.get(0);//TODO use the building num instead
        tvPig.setText("猪只日龄"+bInfo.ageDay+"天");

        for (int nuCot = 0; nuCot < 20; nuCot++) {//loop to set the click event function
            btCottage = (Button) findViewById(cotId[nuCot]);
            if(nuCot < bInfo.unitCount) {
                btCottage.setTextSize(36);
                if(bInfo.unitArrayInfo.get(nuCot).headcount == 0){
                    btCottage.setText("空栏");
                    continue;
                }else{
                    btCottage.setText(bInfo.unitArrayInfo.get(nuCot).unitName);
                }
                coData[nuCot].suggWeight = (double) (bInfo.unitArrayInfo.get(nuCot).suggetFeeding/1000);
                coData[nuCot].finishedWeight = 0;
                coData[nuCot].WeightMark    = 0;

                btCottage.setOnClickListener(new View.OnClickListener() {
                    public void onClick(View v) {
                        int actIdCount;
                        for (actIdCount = 0; actIdCount < 20; actIdCount++) {//loop to find the active card[column]
                            if (cotId[actIdCount] == v.getId()) {
                                break;
                            }
                        }
                        btCottage = (Button) findViewById(cotId[actIdCount]);
                        activeColumn = actIdCount;
                        coData[actIdCount].WeightMark = weightDataBuf.getLastUpdateWeight(); //update the weight

                        buildingInfo bInfo = new buildingInfo();
                        bInfo = farmRunConfig.buildingArrayInfo.get(0);//TODO use the building num instead
                        FarmManager.pigstyInfo pInfo = bInfo.unitArrayInfo.get(actIdCount);
                        weightDataBuf.setAcitveUnit(pInfo.unitID);

                        tvActiveCol.setText("栏号:" + btCottage.getText()+ "\n" //Show the active column number
                                            +"存栏:"+String.valueOf(pInfo.headcount)+"头\n"
                                            +"品种:"+pInfo.pigType);
                        /* bug 20190618 Liz
                        if (coData[actIdCount].WeightMark < coData[actIdCount].suggWeight - coData[actIdCount].finishedWeight) {
                            //fodder is not enought for the cottage, enter into the adding mode
                            String toFeed = new DecimalFormat("#0.00").format(coData[actIdCount].suggWeight - coData[actIdCount].finishedWeight);
                            tvSuggestion.setText("饲料不足，请加料.\n" + "需要饲料" + toFeed + "Kg");//TODO add the data display here
                            isAddingMode = true;
                            return;
                        } else {
                        }*/

                        isAddingMode = false;

                        if (coData[actIdCount].suggWeight - coData[actIdCount].finishedWeight <= 0.05) {
                            v.setBackgroundColor(Color.parseColor("#008080"));//Set the dark color
                        } else {
                            v.setBackgroundColor(Color.parseColor("#F0C042"));//Set the light color
                        }
                    }
                });
            }else{
                btCottage.setVisibility(View.INVISIBLE);//
            }

        }
    }

    public String getAndroidID() {
        String m_szAndroidID = Settings.Secure.getString(getContentResolver(), Settings.Secure.ANDROID_ID);
        return m_szAndroidID;
    }

    /*
     *
     *  Set a timer to reading the weight instrument
     */
    private  void initTimer() {
        Timer timerWeight;
        TimerTask taskWeight;

        final int WHAT = 102;
        final Handler exHandler = new Handler() {//TODO  warning here
            @Override
            public void handleMessage(Message msg) {
                SimpleDateFormat sDateFormat = new SimpleDateFormat("hh:mm:ss");
                switch (msg.what) {
                    case WHAT:
                        //*************Task 1:  Send cmd to weight scaler ********************
                        if (serialScaler.isOpen()) {
                            serialScaler.sendTxt("R");
                        }/* else {
                            Toast.makeText(getBaseContext(), "Error weight scaler tty4 is not open", Toast.LENGTH_SHORT).show();
                        }*/
                        //*************Task 2:  Send weight scaler data to Server ********************
                        if (serialNBiot.isOpen()) {
                            if ((countDebug % 20 == 0)&&(weightDataBuf.isBufferEmpty() == false)) {//T = 300ms * 200 = 1Minute, send data
                                int serIndex = -1;
                                int sizeBuf = 8;//debug only
                                UploadDevDateReq.Builder uploadData = UploadDevDateReq.newBuilder();
                                ReqHeader.Builder reqHeader = ReqHeader.newBuilder();
                                reqHeader.setDevId(getAndroidID());
                                reqHeader.setVersion(1);
                                reqHeader.setTs(System.currentTimeMillis());
                                uploadData.setReqHeader(reqHeader);
                                uploadData.setSerialNo(countDebug/20);
                                for (; weightDataBuf.isBufferEmpty() == false; ) {
                                    serIndex = weightDataBuf.getNextWeightUnitIndex();

                                    ScaleDevRawData.Builder seriBuilder = ScaleDevRawData.newBuilder();
                                    if(weightDataBuf.getActiveColbyIndex(serIndex) == "uploading"){
                                        seriBuilder.setFeedType(LOAD);
                                    }else{
                                        seriBuilder.setFeedType(CHOOSE_PIGSTY);
                                    }
                                    seriBuilder.setPigstyId(weightDataBuf.getActiveColbyIndex(serIndex));
                                    seriBuilder.setTimestamp(weightDataBuf.getRectimelbyIndex(serIndex));
                                    seriBuilder.setCurrentWeight(weightDataBuf.getWeightbyIndex(serIndex));
                                    uploadData.addDevRawData(seriBuilder);
                                }
                                UploadDevDateReq serUpload = uploadData.build();

                                byte[] serBuff = serUpload.toByteArray();
                                sizeBuf += serBuff.length;

                                byte[] head = {(byte) 0xFE, (byte) 0xEF, (byte) 0x9F, (byte) (0xF9),
                                            (byte) 0,//partial?
                                            (byte) (sizeBuf& 0xFF), (byte) (((sizeBuf) >> 8) & 0xFF),//package length?
                                            (byte) 3 };
                                serialNBiot.send(head);
                                serialNBiot.send(serBuff);

                                Toast.makeText(getBaseContext(), "Ser:" + String.valueOf(serIndex)
                                        + "bytes:" + String.valueOf(sizeBuf), Toast.LENGTH_SHORT).show();// debug only
                            }
                        }

                        //*************Task 3:  Send request to Server, to get the config ********************
                        if (serialNBiot.isOpen()) {
                            if (isRequestServerConfig || (countDebug % (150*60)== 8000)) {//T 400*150*60= 1hour
                                int packageLenght = 8;
                                int packageType = 1;


                                ReqHeader.Builder reqHead = ReqHeader.newBuilder();
                                reqHead.setDevId(getAndroidID());
                                reqHead.setTs(System.currentTimeMillis());
                                reqHead.setVersion(0x0601);//TODO
                                reqHead.setConfigVersion(farmRunConfig.getConfigVersion());

                                DevInfoReq.Builder seriBuilder = DevInfoReq.newBuilder();
                                seriBuilder.setReqHeader(reqHead);
                                DevInfoReq serData = seriBuilder.build();
                                byte[] serBuff = serData.toByteArray();
                                packageLenght += serBuff.length;

                                byte[] cmd = {(byte) 0xFE, (byte) 0xEF, (byte) 0x9F, (byte) 0xF9,//head
                                        (byte) 0,//partial?
                                        (byte) (packageLenght & 0xFF), (byte) (((packageLenght) >> 8) & 0xFF),//package length?
                                        (byte) packageType
                                };

                                serialNBiot.send(cmd);
                                serialNBiot.send(serBuff);
                                isRequestServerConfig  = false;
                            }
                        } /*else {
                            Toast.makeText(getBaseContext(), "Error nbiot tty1 is not open",
                                    Toast.LENGTH_SHORT).show();
                        }*/

                        //*************TODO Task 4: Save the feeding data to local file ********************
                        break;
                }
            }
        };
        taskWeight = new TimerTask() {
            @Override
            public void run() {
                Message message = new Message();
                message.what = WHAT;
                message.obj = countDebug++;
                //message.obj = System.currentTimeMillis();
                exHandler.sendMessage(message);
            }
        };

        timerWeight = new Timer();
        timerWeight.schedule(taskWeight, 3000, 400);
    }

    /*
     *
     *  Copy from the demo code by Wangwei, Use 4G/WIFI to reach Server
     */
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
            Toast.makeText(getBaseContext(), "Please,Farmer here ", Toast.LENGTH_SHORT).show();
            tvFarmer.setText(reply.getFarmer());
            return  reply.getFarmer();
        } catch (Exception e) {
            return "Failed... ";
        }
    }
    @Override
    protected void onDestroy() {
        super.onDestroy();
        serialScaler.close();
        serialNBiot.close();
    }
    /*******Get the double value from the string**************/
    private static double getDoubleValue(String str)
    {
        double d = 0;
        if(str != null && str.length() !=0) {
            StringBuffer bf = new StringBuffer();
            char[] chars = str.toCharArray();
            for(int i=0;i<chars.length;i++) {
                char c = chars[i];
                if(c>='0' && c<='9') {
                    bf.append(c);
                } else if(c=='.') {
                    if(bf.length()==0) {
                        continue;
                    } else if(bf.indexOf(".")!=-1) {
                        break;
                    } else {
                        bf.append(c);
                    }
                }else {
                    if(bf.length()!=0) {
                        break;
                    }
                }
            } try {
                d = Double.parseDouble(bf.toString());
            } catch(Exception e) {}
        }
        return d;
    }

    /*
     *   New a thread to wating the data from weight scaler
     */
    private void initWeightScaler() {

        serialScaler = new SerialHelper("/dev/ttyS4", 9600) {
            @Override
            protected void onDataReceived(final ComBean comBean) {
                runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        Double weight = 0.00;
                        ColumnData columnData;
                        //Debug start
                        logListAdapter.addData(comBean.sRecTime + ":   " + (new String(comBean.bRec)));//Print the sting intstead of HEX
                        receDisp.smoothScrollToPosition(logListAdapter.getData().size());
                        //Debug end
                        weight = getDoubleValue(new String(comBean.bRec));

                        if (isAddingMode) {
                            weightDataBuf.setAcitveUnit("uploading");
                            weightDataBuf.addWeight(System.currentTimeMillis(), weight);//use -1 to indicate adding mode
                            tvSuggestion.setText("上料状态\n"+"饲料"+String.valueOf(weight)+"kg");
                            return; // in adding fodder mode, nothing to do
                        }

                        /*****Write data to ring buffer***********/
                        weightDataBuf.addWeight(System.currentTimeMillis(), weight);
                        // if (weightDataBuf.addWeight(comBean.sRecTime, weight))
                        {
                            //weight changed, get the active cottage
                            btCottage = (Button) findViewById(cotId[activeColumn]);
                            columnData = coData[activeColumn];
                            columnData.finishedWeight += (columnData.WeightMark - weight);
                            columnData.WeightMark = weight;//update the water mark

                            if (columnData.suggWeight - columnData.finishedWeight <= 0.1) {
                                btCottage.setBackgroundColor(Color.parseColor("#008080"));//Set the dark color
                            } else {
                                btCottage.setBackgroundColor(Color.parseColor("#F0C042"));//Set the light color
                            }

                            String suggTemp = new DecimalFormat("#00.00").format(columnData.suggWeight);
                            String finishedfeedTemp = new DecimalFormat("#00.00").format(columnData.finishedWeight);
                            String tofeedTemp = new DecimalFormat("#00.00").format(columnData.suggWeight - columnData.finishedWeight);
                            tvSuggestion.setText(
                                    "推荐:    " + suggTemp + "kg\n"
                                            + "已喂:    " + finishedfeedTemp + "kg\n"
                                            + "还剩:    " + tofeedTemp+"kg");
                        }
                    }
                });
            }
        };
        try {
            serialScaler.open();
            Toast.makeText(getBaseContext(), "Connect to the weight scaler", Toast.LENGTH_SHORT).show();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    /*
     *   1,Load data from the latest config file in local directory;
     *   2,if not exist, just wait;
     *   3,request cmd is send on every day;
     *   4,A thread is waiting the data from Server by NBiot
     */
    private  void initNBiot(){

        final FarmManager farmConfig;

        farmConfig = new FarmManager();
        serialNBiot = new SerialHelper("/dev/ttyS1", 9600) {
            @Override
            protected void onDataReceived(final ComBean comBean) {
                runOnUiThread(new Runnable() {
                    @Override
                    public void run() {
                        /*
                         * Section 1: Byte[0~3], Sync Head, 0xFE 0xEF 0x9F 0xF9
                         * Section 2: Byte[4], package entity type 0:full 1:Start 2:Mid 3:End
                         * Section 3: Byte[5~6], package length
                         * Section 4: Byte[7], package cmd type 1: request 2:return's config
                         * Section 5: Byte[8~], data area
                         */
                        //TODO check the head, and package type
                        byte[]rBuf = comBean.bRec;
                        SaveByestoFile("AFS20190616.cfg",rBuf,8,rBuf.length-8);
                        tvFarmer.setText("自动重启中");
                        try {
                            Thread.sleep(3000);
                        } catch (InterruptedException e) {
                            e.printStackTrace();
                        }
                        Intent intent = getBaseContext().getPackageManager().getLaunchIntentForPackage(getBaseContext().getPackageName());
                        intent.addFlags(Intent.FLAG_ACTIVITY_CLEAR_TOP);
                        intent.putExtra("REBOOT","reboot");
                        startActivity(intent);
                        //Toast.makeText(getBaseContext(), "Data from NBiot: "+new String(comBean.bRec), Toast.LENGTH_SHORT).show();
                    }
                });
            }
        };
        try {
            serialNBiot.open();
            Toast.makeText(getBaseContext(), "Connect to the NBiot module", Toast.LENGTH_SHORT).show();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    /*
     *   Register the  click events functions
     *   Open the UART  to weight measurement
     */
    private void initClickFunc() {
        /**********Open the UART debug only *************/
        btOpenPort.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                try {
                    serialScaler.close();
                    if (radioGroup.getCheckedRadioButtonId() == R.id.radioButton1){ //Sent text
                        serialScaler.setPort("/dev/ttyS1");
                        Toast.makeText(getBaseContext(), "Select the com[1]", Toast.LENGTH_SHORT).show();
                    }else{
                        serialScaler.setPort("/dev/ttyS4");
                        Toast.makeText(getBaseContext(), "Select the com[4]", Toast.LENGTH_SHORT).show();
                    }

                    serialScaler.open();
                    Toast.makeText(getBaseContext(), "In debug mode:Open uart once you change the com selection", Toast.LENGTH_SHORT).show();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        });


        /**********Send command to the electronic scale debug only*************/
        btSenddata.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if (edInputdata.getText().toString().length() > 0) {
                    if (serialScaler.isOpen()) {
                        serialScaler.sendTxt(edInputdata.getText().toString());
                    } else {
                        Toast.makeText(getBaseContext(), "open the com port", Toast.LENGTH_SHORT).show();
                    }
                } else {
                    Toast.makeText(getBaseContext(), "Fill the data to send", Toast.LENGTH_SHORT).show();
                }
            }
        });

        /**********adjust the electronic scale to Zero*************/
        btZero.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if (serialScaler.isOpen()) {
                    serialScaler.sendTxt("T");
                } else {
                    Toast.makeText(getBaseContext(), "open the com port", Toast.LENGTH_SHORT).show();
                }
            }
        });
        /******send cmd to read the weight instrument
        btMeas.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if (serialScaler.isOpen()) {
                    serialScaler.sendTxt("R");
                } else {
                    Toast.makeText(getBaseContext(), "open the com port", Toast.LENGTH_SHORT).show();
                }
            }
        });
         ******/
        /***********Uploading *******************/
        btUploading.setOnClickListener((new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                isAddingMode = true;
            }
        }));
        /*
        btExit.setOnClickListener(new View.OnClickListener(){
            @Override
            public void onClick(View view){
                isAddingMode = true;
                serialScaler.close();
                // TODO sent data to sever before exit
                android.os.Process.killProcess(android.os.Process.myPid());
            }
        });
        */
    }
}
