package com.serialcom;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.List;
import java.util.ServiceConfigurationError;

import cn.iopig.feed.scale.grpc.PigHouseInfo;
import cn.iopig.feed.scale.grpc.PigstyInfo;
import cn.iopig.feed.scale.grpc.PigstyInfoRes;
import cn.iopig.feed.scale.grpc.ReqHeader;

public class FarmManager{
    public FarmManager(){
        farmer = null;
        address = null;
        serviceline = null;
        farmerID = null;
        configVersion = 0;
    }
    private String farmer;
    private String farmerID;
    private String address;
    private String serviceline;
    private int cottageCount;
    private int configVersion;
    class AFSInfo{
        int AFSVersion=0;
        String DeviceID=null;
        int timestamp=0;
    }
    AFSInfo headInfo;

    class pigstyInfo{
        String unitID;
        String unitName;
        String pigType;
        int headcount;
        int averageWeight;
        int suggetFeeding;
        long lastFeedingTimeStamp;
        String []pigID;
    }
    static class buildingInfo{
        String buildingId;
        int unitCount;
        String ageDay;
        //pigstyInfo[] unitInfo;
        ArrayList<pigstyInfo> unitArrayInfo;
    }
    ArrayList<buildingInfo> buildingArrayInfo;

    public String getFarmer(){return farmer;}
    public int getConfigVersion(){return configVersion;}
    public boolean updateFarmInfo(PigstyInfoRes configFromServer){
        //TODO init the head?
        farmer      = configFromServer.getFarmer();
        farmerID    = configFromServer.getPigFarmId();
        ReqHeader.Builder reqHead = configFromServer.getReqHeader().newBuilder();
        configVersion = reqHead.getConfigVersion();

        buildingArrayInfo = new ArrayList<buildingInfo>();

        List<PigHouseInfo> houseInfoList = configFromServer.getPigHouseInfoList();
        for (PigHouseInfo houseInfo : houseInfoList){
            buildingInfo bInfo = new buildingInfo();
            bInfo.buildingId = houseInfo.getHouseId();
            bInfo.unitCount = houseInfo.getPigstyInfoCount();
            bInfo.ageDay    = houseInfo.getAGE();
            bInfo.unitArrayInfo = new ArrayList<pigstyInfo>();

            List<PigstyInfo>pigstyList = houseInfo.getPigstyInfoList();
            for(PigstyInfo pigtyInfo:pigstyList){
                pigstyInfo pigstyconfig = new pigstyInfo();
                pigstyconfig.averageWeight  = pigtyInfo.getAverageWeight();
                pigstyconfig.headcount      = pigtyInfo.getPigNum();
                pigstyconfig.lastFeedingTimeStamp = pigtyInfo.getLastFed();
                pigstyconfig.unitID         = pigtyInfo.getPigstyId();
                pigstyconfig.unitName       = pigtyInfo.getStyName();
                pigstyconfig.pigType        = pigtyInfo.getPigSpecies();
                pigstyconfig.suggetFeeding  = pigtyInfo.getAdviseWeight();

                List<String>pigIDList = pigtyInfo.getPigIdList();
                int index=0;
                pigstyconfig.pigID = new String[pigIDList.size()];
                for(String pigID:pigIDList){
                    pigstyconfig.pigID[index] = pigIDList.get(index);
                    index++;
                }
                bInfo.unitArrayInfo.add(pigstyconfig);
            }
            buildingArrayInfo.add(bInfo);
        }
        return true;
    }
}
