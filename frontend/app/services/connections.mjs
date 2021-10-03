import { server } from "../config.mjs";

async function SupportService(value){
    return fetch(server + "support/" + value).then(res => res.json());
}

async function ResistanceService(value){
    return fetch(server + "resistance/" + value).then(res => res.json());
}

async function LoadPairService(pair){
    return fetch(server + "load/" + pair).then(res => res.json());
}

async function GetPairService(){
    return fetch(server + "get").then(res => res.json());
}

async function EmaService(){
    return fetch(server + "emas").then(res => res.json());
}

async function BacktrackService(days){
    return fetch(server + "backtrack/" + days);
}

async function GetRSIService(){
    return fetch(server + "rsi").then(res => res.json());
}

async function GetWedgeService(){
    return fetch(server + "wedge").then(res => res.json());
}

async function GetTrendsService(){
    return fetch(server + "trends").then(res => res.json());
}

async function GetTargetsService(){
    return fetch(server + "targets").then(res => res.json());
}

export { SupportService, EmaService, ResistanceService, LoadPairService, GetPairService, BacktrackService, GetRSIService, GetWedgeService, GetTrendsService, GetTargetsService }