$(document).ready(function(){
    var BASE_URL = "http://localhost:8080";

   $("#btn-add-app").click(function() {
       var name = $("#ipt-app-name").val();
       $.post(BASE_URL + "/dashboard/add", {
           "name": name
       },
       function(ret) {
           if(ret.errno == 0) {
              console.log("添加成功");
              return;
           }


       })
   });
});
