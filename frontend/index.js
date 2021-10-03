var ChartServices = {
    backtrack: () => {
        import('./app/app.mjs').then(module => {
            module.Backtrack(document.getElementById('backtrack_days').value);
        });
    },
    newline: () => {
        import('./app/app.mjs').then(module => {
            module.NewLine();
        });
    },
    newsequence: () => {
        import('./app/app.mjs').then(module => {
            module.NewSequence();
        });
    },
    stopsequence: () => {
        import('./app/app.mjs').then(module => {
            module.StopSequence();
        });
    },
    deletelastline: () => {
        import('./app/app.mjs').then(module => {
            module.DeleteLastLine();
        });
    },
    openPair: () => {
        var pair = document.getElementById("pair").value;
        import('./app/app.mjs').then(module => {
            module.Start(pair, LightweightCharts);
        });
    }
}