import { GetRSIService } from "../services/connections.mjs";
import { DrawInfiniteLine } from "./lines.mjs";

function LoadRSI(rsiChart, rsiSeries, cb){
    rsiChart.removeSeries(rsiSeries);
    rsiSeries = rsiChart.addLineSeries();

    GetRSIService()
    .then(data => {
        var rsiData = [];
        var dayZero = new Date();
        dayZero.setDate(dayZero.getDate() - data.rsi.length);

        for(var i = 0; i < data.rsi.length; i++){
            rsiData.push({
                time: new Date(dayZero.setDate(dayZero.getDate() + 1)).getTime()/1000,
                value: data.rsi[i]
            });
        }

        rsiSeries.setData(rsiData);
        DrawInfiniteLine(rsiSeries, 70, "largeDashed");
        DrawInfiniteLine(rsiSeries, 30, "largeDashed");

        cb(rsiChart, rsiSeries);
    })

    .catch(err => console.log(err));
}


export { LoadRSI }