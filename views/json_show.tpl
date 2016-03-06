<!DOCTYPE html>

<html>
<head>
 <meta charset="UTF-8">
  <title>技术人员价值评测系统</title>
  <base target=_blank>
<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
</head>


<div id="treeview" class="col-xs-3"></div>

<div class="col-lg-9">
  <table class="table table-striped">
    <thead>
      <tr>
        <th>#序号</th>
        <th>单位</th>
        <th>分院</th>
        <th>科室</th>
        <th>价值分类</th>
        <th>价值项</th>
        <th>选择项</th>
        <th>操作</th>
      </tr>
    </thead>
  </table>
</div>

<script type="text/javascript">
$(function() {
         // $('#treeview').treeview('collapseAll', { silent: true });
          $('#treeview').treeview({
          data: [{{.json}}],//defaultData,
          // data:alternateData,
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