const chartProperties = {
    width: 1500, 
    height: 650,
    timeScale: {
        timeVisible: true,
        secondsVisible: false,
    },
    crosshair: {
		mode: LightweightCharts.CrosshairMode.Normal,
	},
	timeScale: {
		borderColor: 'rgba(197, 203, 206, 1)',
	},
	handleScroll: {
		vertTouchDrag: false,
	}
}

const rsiProperties = {
	width: 1500,
	height: 120,
	crosshair: {
		mode: LightweightCharts.CrosshairMode.Normal,
	},
	timeScale: {
		visible: false,
	},
}

const server = "http://localhost:8080/";

export { chartProperties, server, rsiProperties }