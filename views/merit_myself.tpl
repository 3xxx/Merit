<!-- 展示个人的价值列表：已通过，待提交 管理员查看：已通过……-->
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Merit价值管理系统</title>

<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
<script src="/static/js/bootstrap-treeview.js"></script>
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

<div class="col-lg-12">
  <h2>{{.UserNickname}}</h2>
  <ul class="nav nav-tabs">
    <li class="active"><a href="#employee" data-toggle="tab">已通过</a></li>
    <li><a href="#year" data-toggle="tab">待提交</a></li>
    <li><a href="#proj" data-toggle="tab">分布</a></li>
  </ul>

  <div class="tab-content">
    <div class="tab-pane fade in active" id="employee">
    <br>
      <div class="form-inline">
        <table class="table table-striped">
          <thead>
            <tr>
              <th>#</th>
              <th>Title</th>
              <th>choose</th>
              <th>mark</th>
              <th>ParentTitle</th>
              <th>操作</th>
            </tr>
          </thead>
          <tbody>
            {{range $k,$v :=.topics}}
            <tr>
              <td>{{$k|indexaddone}}</td>
              <td>{{.Title}}</td>
              <td>{{.Choose}}</td>
              <td>{{.Mark}}</td>
              {{range $k1,$v1 :=$.category}}
              {{if eq $v.ParentId $v1.Id}}
              <td>{{.Title}}</td>
              {{end}}
              {{end}}
              <td>
               <a href="/view?id={{.Id}}"><i class="glyphicon glyphicon-open"></i>详细</a>
                <a href="/modify?id={{.Id}}"><i class="glyphicon glyphicon-edit"></i>修改</a>
                <a href="/delete?id={{.Id}}"><i id="delete" class="glyphicon      glyphicon-remove-sign"></i>删除</a>
              </td>  
            </tr>
            {{end}}
          </tbody>
        </table>
      </div>
    </div>

    <div class="tab-pane fade" id="year">
      <p>这里将显示待提交的价值记录。</p>
    </div>

    <div class="tab-pane fade" id="proj">
      <p>这里将显示价值分布。</p>
    </div>

  </div>
</div>
</body>
</html>