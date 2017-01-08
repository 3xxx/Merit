<!-- 展示个人的价值列表：已通过，待提交 管理员查看：已通过……-->
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>Merit价值管理系统</title>
      <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
      <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
      <script src="/static/js/bootstrap-treeview.js"></script>
      <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
      <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
    
      <script type="text/javascript" src="/static/js/moment.min.js"></script>
      <script type="text/javascript" src="/static/js/daterangepicker.js"></script>
      <link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css"/>
    
      <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
      <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
      <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
    
      <script src="/static/js/moment-with-locales.min.js"></script>
      <script type="text/javascript" src="/static/js/echarts.min.js"></script>
    
      <script src="/static/js/bootstrap-table-filter-control.js"></script>
      <script src="/static/js/bootstrap-table-export.min.js"></script>
      <script src="/static/js/tableExport.js"></script>
      <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
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
    <h2>{{.Sectitle}}</h2>
    <ul class="nav nav-tabs">
      <li class="active"><a href="#employee" data-toggle="tab">已通过</a></li>
      <li><a href="#year" data-toggle="tab">待提交</a></li>
      <li><a href="#proj" data-toggle="tab">分布</a></li>
    </ul>

    <div class="tab-content">
      <div class="tab-pane fade in active" id="employee">
      <br>
        <div class="form-inline">
            <table id="table"
                  data-toggle="table"
                  data-url="/merit/secofficedata"
                  data-search="true"
                  data-show-refresh="true"
                  data-show-toggle="true"
                  data-show-columns="true"
                  data-show-export="true"
                  data-query-params="queryParams"
                  >
              <thead>        
              <tr>
                  <th data-formatter="index1">#</th>
                  <th data-field="Name" data-sortable="true">姓名</th>
                  <th data-field="Count">项数</th>
                  <th data-field="Sigma" data-sortable="true">汇总</th>
                  <th data-field="action" data-formatter="actionFormatter1">详细</th>
                </tr>
              </thead>
            </table>
        <!-- <table class="table table-striped">
          <thead>
            <tr>
              <th><span style="cursor: pointer">#</span></th>
              <th><span style="cursor: pointer">Name</span></th>
              <th><span style="cursor: pointer">Numbers</span></th>
              <th><span style="cursor: pointer">Marks</span></th>
              <th><span style="cursor: pointer">Secoffic</span></th>
              <th><span style="cursor: pointer">Department</span></th>
              <th>操作</th>
            </tr>
          </thead>
          <tbody>
            {{range $k,$v :=.Employee}}
            <tr>
              <td>{{$k|indexaddone}}</td>
              <td>{{.Name}}</td>
              <td>{{.Numbers}}</td>
              <td>{{.Marks}}</td>
              <td>{{.Keshi}}</td>
              <td>{{.Department}}</td>
              <td>
               <a href="/merit/myself?id={{.Id}}&level=3"><i class="glyphicon glyphicon-open"></i>详细</a>
              </td>  
            </tr>
            {{end}}
          </tbody>
        </table> -->
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