<!-- iframe里展示个人详细情况-->
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <title>情况汇总</title>
  <!-- <base target=_blank> -->
<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
 <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script> 
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
<!-- <style type="text/css">
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
</script> -->
</head>


<!-- <div id="treeview" class="col-xs-3"></div> -->

<div class="col-lg-12">
  <table class="table table-striped">
    <thead>
      <tr>
        <th>#</th>
        <th>项目编号</th>
        <th>项目名称</th>
        <th>项目阶段</th>
        <th>成果编号</th>
        <th>成果名称</th>
        <th>成果类型</th>
        <th>成果计量单位</th>
        <th>成果数量</th>
        <th>编制、绘制</th>
        <th>设计</th>
        <th>校核</th>
        <th>审查</th>
      </tr>
    </thead>

    <tbody>
      <tr><th colspan=13>图纸</th></tr>
      {{range $k,$v :=.Catalogtuzhi}}
      <tr>
        <th>{{$k}}</th>
        <th>{{.ProjectNumber}}</th>
        <th>{{.ProjectName}}</th>
        <th>{{.DesignStage}}</th>

        <th>{{.Tnumber}}</th>
        <th>{{.Name}}</th>
        <th>{{.Category }}</th>
        <th>{{.Page}}</th>
        <th>{{.Count }}</th>
        <th>{{.Drawn }}</th>
        <th>{{.Designd}}</th>
        <th>{{.Checked}}</th>
         <th>{{.Examined}}</th> 
      </tr>
      {{end}}

<tr><th colspan=13>报告</th></tr>
    {{range $k,$v :=.Catalogbaogao}}
      <tr>
        <th>{{$k}}</th>
        <th>{{.ProjectNumber}}</th>
        <th>{{.ProjectName}}</th>
        <th>{{.DesignStage}}</th>

        <th>{{.Tnumber}}</th>
        <th>{{.Name}}</th>
        <th>{{.Category }}</th>
        <th>{{.Page}}</th>
        <th>{{.Count}}</th>
        <th>{{.Drawn}}</th>
        <th>{{.Designd}}</th>
        <th>{{.Checked}}</th>
        <th>{{.Examined}}</th>  
      </tr>
      {{end}}

<tr><th colspan=13>计算书</th></tr>
      {{range $k,$v :=.Catalogjisuanshu}}
      <tr>
        <th>{{$k}}</th>
        <th>{{.ProjectNumber}}</th>
        <th>{{.ProjectName}}</th>
        <th>{{.DesignStage}}</th>

        <th>{{.Tnumber}}</th>
        <th>{{.Name}}</th>
        <th>{{.Category }}</th>
        <th>{{.Page}}</th>
        <th>{{.Count}}</th>
        <th>{{.Drawn}}</th>
        <th>{{.Designd}}</th>
        <th>{{.Checked}}</th>
        <th>{{.Examined}}</th>  
      </tr>
      {{end}}
 
 <tr><th colspan=13>修改单</th></tr>
      {{range $k,$v :=.Catalogxiugaidan}}
      <tr>
        <th>{{$k}}</th>
        <th>{{.ProjectNumber}}</th>
        <th>{{.ProjectName}}</th>
        <th>{{.DesignStage}}</th>

        <th>{{.Tnumber}}</th>
        <th>{{.Name}}</th>
        <th>{{.Category }}</th>
        <th>{{.Page}}</th>
        <th>{{.Count}}</th>
        <th>{{.Drawn}}</th>
        <th>{{.Designd}}</th>
        <th>{{.Checked}}</th>
        <th>{{.Examined}}</th>  
      </tr>
      {{end}}

<tr><th colspan=13>大纲</th></tr>
      {{range $k,$v :=.Catalogdagang}}
      <tr>
        <th>{{$k}}</th>
        <th>{{.ProjectNumber}}</th>
        <th>{{.ProjectName}}</th>
        <th>{{.DesignStage}}</th>

        <th>{{.Tnumber}}</th>
        <th>{{.Name}}</th>
        <th>{{.Category }}</th>
        <th>{{.Page}}</th>
        <th>{{.Count}}</th>
        <th>{{.Drawn}}</th>
        <th>{{.Designd}}</th>
        <th>{{.Checked}}</th>
        <th>{{.Examined}}</th>  
      </tr>
      {{end}}

<tr><th colspan=13>标书</th></tr>
      {{range $k,$v :=.Catalogbiaoshu}}
      <tr>
        <td>{{$k}}</td>
        <td>{{.ProjectNumber}}</td>
        <td>{{.ProjectName}}</td>
        <td>{{.DesignStage}}</td>

        <td>{{.Tnumber}}</td>
        <td>{{.Name}}</td>
        <td>{{.Category }}</td>
        <td>{{.Page}}</td>
        <td>{{.Count }}</td>
        <td>{{.Drawn }}</td>
        <td>{{.Designd}}</td>
        <td>{{.Checked}}</td>
        <td>{{.Examined}}</td>  
      </tr>
      {{end}}
    </tbody>

  </table>
</div>

<script type="text/javascript">
// $(function() {
         // $('#treeview').treeview('collapseAll', { silent: true });
          // $('#treeview').treeview({
          // data: [{{.json}}],//defaultData,
          // data:alternateData,
          // levels: 5,// expanded to 5 levels
          // enableLinks:true,
          // showTags:true,
          // collapseIcon:"glyphicon glyphicon-chevron-up",
          // expandIcon:"glyphicon glyphicon-chevron-down",
//         });
// });


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