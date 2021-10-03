import { EmaService } from "../services/connections.mjs";


function LoadVolume(chart, chData){
    var volumeSeries = chart.addHistogramSeries({
        color: '#26a69a',
        priceFormat: {
            type: 'volume',
        },
        priceScaleId: '',
        scaleMargins: {
            top: 0.8,
            bottom: 0,
        },
    });

    var result = [];
    for(var i = 0; i < chData.length; i++)
    {
        result.push(
            {
                time: chData[i].time, 
                value: chData[i].volume,
                color: ( chData[i].close > chData[i].open ? "#00cc66" : "#ff5050" )
            }
        );
    }
    volumeSeries.setData(result);
}

function LoadEmas(chart, chData, cb){
    var ema56Line, ema200Line, ema500Line;
    EmaService()
    .then(data => {

        var result56 = [];
        var result200 = [];
        var result500 = [];

        for(var i = 0; i < data.ema56.length; i++)
        {
            result56.push({
                time: chData[i].time,
                value: data.ema56[i]
            });

            result200.push({
                time: chData[i].time,
                value: data.ema200[i]
            });

            result500.push({
                time: chData[i].time,
                value: data.ema500[i]
            });
        }

        ema56Line = chart.addLineSeries({
            color: 'blue',
            lineWidth: 1
        });

        ema200Line = chart.addLineSeries({
            color: 'purple',
            lineWidth: 2
        });

        ema500Line = chart.addLineSeries({
            color: 'orange',
            lineWidth: 3
        });

        ema56Line.setData(result56);
        ema200Line.setData(result200);
        ema500Line.setData(result500);

        cb(ema56Line, ema200Line, ema500Line);
    })
    .catch(err => console.log(err))
}

export { LoadVolume, LoadEmas }