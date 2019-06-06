package com.serialcom;


import com.chad.library.adapter.base.BaseQuickAdapter;
import com.chad.library.adapter.base.BaseViewHolder;
import com.serialcom.R;

import java.util.List;

/**
 * Author
 * Created Time 2017/12/14.
 */
public class LogListAdapter extends BaseQuickAdapter<String, BaseViewHolder> {

    public LogListAdapter(List list) {
        super(R.layout.item_layout, list);
    }

    @Override
    protected void convert(BaseViewHolder helper, String item) {

        helper.setText(R.id.textView, item);

    }

}
