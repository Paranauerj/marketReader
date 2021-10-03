import { GetTrendsService } from "../services/connections.mjs";

function GetTrends(domShort, domMid, domLong){
    GetTrendsService()
    .then(data => {
        fillDomElement(domShort, data.short_term);
        fillDomElement(domMid, data.mid_term);
        fillDomElement(domLong, data.long_term);
    });
}

function fillDomElement(dom, arr){
    if(arr){
        for(var i = 0; i < arr.length; i++){
            dom.innerHTML += "<br/>" + arr[i];
        }
    }
}

export { GetTrends }