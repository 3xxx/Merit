<!-- 展示科室总体情况-->
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <title>科室情况汇总</title>
  <!-- <base target=_blank> -->
<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
 <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script> 
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
<style type="text/css">
a:active{text:expression(target="_blank");}
i#delete
{
color:#DC143C;
}
</style>
<script type="text/javascript">
  var allLinks=document.getElementsByTagName("a");
for(var i=0;i!=allLinks.length; i++){
allLinks[i].target="_blank";
}

</script>
</head>


<!-- <div id="treeview" class="col-xs-3"></div> -->

<div class="col-lg-12">
  <table class="table table-striped">
    <thead>
      <tr>
        <th>#</th>
        <th>姓名</th>
        <th>编制</th>
        <th>设计</th>
        <th>校核</th>
        <th>审查</th>
        <th>汇总</th>
        <th>详细</th>
      </tr>
    </thead>

    <tbody>
      {{range $k,$v :=.Employee}}
      <tr>
        <th>{{$k}}</th>
        <th>{{.Name}}</th>
        <th>{{.Drawn}}</th>
        <th>{{.Designd}}</th>
        <th>{{.Checked}}</th>
        <th>{{.Examined}}</th>
        <th>{{.Sigma}}</th>
        <th>
         <a href="/secofficeshow?secid={{.Id}}&level=3"><i class="glyphicon glyphicon-open"></i>详细</a>
        </th>  
      </tr>
      {{end}}
    </tbody>

  </table>
</div>

<script type="text/javascript">
$(function() {
         // $('#treeview').treeview('collapseAll', { silent: true });
          $('#treeview').treeview({
          data: [{{.json}}],//defaultData,
          // data:alternateData,
          levels: 5,// expanded to 5 levels
          enableLinks:true,
          showTags:true,
          // collapseIcon:"glyphicon glyphicon-chevron-up",
          // expandIcon:"glyphicon glyphicon-chevron-down",
        });
});


  $(document).ready(function() {
  $("table").tablesorter({sortList: [[6,1]]});
  // $("#ajax-append").click(function() {
  //    $.get("assets/ajax-content.html", function(html) {
  //     // append the "ajax'd" data to the table body
  //     $("table tbody").append(html);
  //     // let the plugin know that we made a update
  //     $("table").trigger("update");
  //     // set sorting column and direction, this will sort on the first and third column
  //     var sorting = [[2,1],[0,0]];
  //     // sort on the first column
  //     $("table").trigger("sorton",[sorting]);
  //   });
  //   return false;
  // });
});
</script>
</body>
</html>
<!-- <button type="button" class="btn btn-primary btn-lg" style="color: rgb(212, 106, 64);">
<span class="glyphicon glyphicon-user"></span> User
</button>

<button type="button" class="btn btn-primary btn-lg" style="text-shadow: black 5px 3px 3px;">
<span class="glyphicon glyphicon-user"></span> User
</button> -->