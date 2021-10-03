import { DrawLineSeries } from "../chart/lines.mjs";
import { DrawLine } from "../chart/lines.mjs";
import { GetTargetsService } from "../services/connections.mjs";

function DrawTargets(chart, chData, cb){
    GetTargetsService()
    .then(data => {
        var resultTarget1 = [];
        var  resultTarget2 = [];
        var resultTarget1Line, resultTarget2Line;

        for(var i = 0; i < data.tops.length; i++)
        {
            resultTarget1.push({
                time: chData[i].time,
                value: data.tops[i].value
            });

            resultTarget2.push({
                time: chData[i].time,
                value: data.lows[i].value
            });
        }

        resultTarget1Line = chart.addLineSeries({
            color: 'green',
            lineWidth: 2,
            lineStyle: LightweightCharts.LineStyle.Dotted
        });

        resultTarget2Line = chart.addLineSeries({
            color: 'red',
            lineWidth: 2,
            lineStyle: LightweightCharts.LineStyle.Dotted
        });

        resultTarget1Line.setData(resultTarget1);
        resultTarget2Line.setData(resultTarget2);

        cb(resultTarget1Line, resultTarget2Line);
    })
    .catch(err => console.log(err))
}

export { DrawTargets }