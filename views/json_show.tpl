<!-- 展示所有侧栏的界面 将来修改为管理员目录-->
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <title>技术人员价值评测系统</title>
  <!-- <base target=_blank> -->
<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
<style type="text/css">
a:active{text:expression(target="_blank");}
</style>
<script type="text/javascript">
  var allLinks=document.getElementsByTagName("a");
for(var i=0;i!=allLinks.length; i++){
allLinks[i].target="_blank";
}
</script>
</head>


<div id="treeview" class="col-xs-3"></div>

<div class="col-lg-9">
  <table class="table table-striped">
    <thead>
      <tr>
        <th>#</th>
        <th>Title</th>
        <th>choose</th>
        <th>mark</th>
        <th>content</th>
        <th>操作</th>
      </tr>
    </thead>
    <tbody>
      {{range $k,$v :=.topics}}
      <tr>
        <th>{{$k}}</th>
        <th>{{.Title}}</th>
                   <th>{{.Choose}}</th>
                   <th>{{.Mark}}</th>
                   <th>{{.Content}}</th>
                    <th>
                  <a href="/">显示成果</a>
                  <a href="/">修改</a>
                  <a href="/">删除</a>
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
</script>
</body>
</html>
<!-- <button type="button" class="btn btn-primary btn-lg" style="color: rgb(212, 106, 64);">
<span class="glyphicon glyphicon-user"></span> User
</button>

<button type="button" class="btn btn-primary btn-lg" style="text-shadow: black 5px 3px 3px;">
<span class="glyphicon glyphicon-user"></span> User
</button> -->