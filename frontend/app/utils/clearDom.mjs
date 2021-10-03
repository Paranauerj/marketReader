function ClearDoms(doms){
    for(var i = 0; i < doms.length; i++){
        doms[i].innerHTML = "";
    }
}

export { ClearDoms }