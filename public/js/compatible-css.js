$(document).ready(function(){
    var pageLogo = $(".page-logo");
    var menuSearchBtn = $("#menu-search-btn");//搜索框的搜索按钮

    if(navigator.userAgent.indexOf("MSIE 7.0") > 0){
        menuSearchBtn.css({
            top:6,
            height:19,
            padding: "3px 0 0 10px"
        });
    }

    if(navigator.userAgent.indexOf("MSIE 7.0") < 0 && navigator.userAgent.indexOf("MSIE 6.0") < 0  && navigator.userAgent.indexOf("MSIE 8.0") < 0){
        pageLogo.corner("5px bottom");
    }

    //alert(navigator.userAgent);
});
