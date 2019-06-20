package com.serialcom;

public class WeightBuffer {
    public WeightBuffer(int bufferSize) {
        this.stampIndex = 0;
        this.uploadIndex = 0;
        this.bufferSize = bufferSize;
        this.weightbuf= new WeightUnit[this.bufferSize];
        for (int count = 0; count < bufferSize; count++){
            this.weightbuf[count] = new WeightUnit();
        }
        this.acitveUnit = null;
    }

    class WeightUnit{
        String activeColumn = null;
        long   recTime = 0;//ms
        double weightFodder = 0.0;
    }
    private int stampIndex;
    private int uploadIndex;
    private int bufferSize;
    private String acitveUnit;
    WeightUnit[] weightbuf;

    public void setAcitveUnit(String strActiveUnit)
    {
        acitveUnit = strActiveUnit;
    }

    public boolean isBufferEmpty(){
        if (uploadIndex == stampIndex)
            return true;
        else
            return false;
    }

    public int getNextWeightUnitIndex(){
        int index = uploadIndex;
        uploadIndex++;
        if(uploadIndex ==bufferSize)
           uploadIndex = 0;
        return index;
    }

    public String getActiveColbyIndex(int index){
        return weightbuf[index].activeColumn;
    }
    public long getRectimelbyIndex(int index){
       return weightbuf[index].recTime;
    }
    public double getWeightbyIndex(int index){
        return weightbuf[index].weightFodder;
    }

    public boolean addWeight(long recTime, double weightFodder){
        int lastIndex;
        if (stampIndex == 0)
            lastIndex = bufferSize -1;
        else
            lastIndex = stampIndex - 1;

        if (((weightbuf[lastIndex].weightFodder - weightFodder) < 0.01)
                && ((weightbuf[lastIndex].weightFodder - weightFodder) > -0.01))
        {
            return false;// minor change, just return;
        }
        weightbuf[stampIndex].activeColumn = acitveUnit;
        weightbuf[stampIndex].weightFodder = weightFodder;
        weightbuf[stampIndex].recTime = recTime;

        stampIndex++;
        if (stampIndex == bufferSize){//wrap around
            stampIndex = 0;
        }
        return true;
    }
    public double getLastUpdateWeight(){
        int index = stampIndex;
            if(index == 0)
                index=bufferSize-1;
            else
                index--;
        return weightbuf[index].weightFodder;
    }
}
