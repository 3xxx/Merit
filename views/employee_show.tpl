<!-- iframe里展示个人详细情况，这个是自己查看显示处理按钮-->
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <title>情况汇总</title>

<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
 <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
 <script src="/static/js/bootstrap-treeview.js"></script>
 <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script> 
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
<script type="text/javascript" src="/static/js/moment.min.js"></script>
  <script type="text/javascript" src="/static/js/daterangepicker.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css" />

 <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
<script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script> 

<script src="/static/js/moment-with-locales.min.js"></script> 
</head>

<!-- <div class="form-group"> -->
<div class="col-lg-12">
<h2>{{.UserNickname}}</h2>
<div>
<!-- <form class="form-inline" method="get" action="/secofficeshow" enctype="multipart/form-data"> -->
<div class="form-inline">
  <input type="hidden" id="secid" name="secid" value="{{.Secid}}"/>
  <input type="hidden" id="level" name="level" value="{{.Level}}"/>
  <input type="hidden" id="key" name="key" value="modify"/>
  <div class="form-group">
    <label for="taskNote">统计周期：</label>
    <input type="text" class="form-control" name="datefilter" id="datefilter" value="" placeholder="选择时间段(默认最近一个月)"/>
  </div>
  <script type="text/javascript">
$(function() {
  $('input[name="datefilter"]').daterangepicker({
      autoUpdateInput: false,
      locale: {
          cancelLabel: 'Clear'
      }
  });
  $('input[name="datefilter"]').on('apply.daterangepicker', function(ev, picker) {
      $(this).val(picker.startDate.format('YYYY-MM-DD') + ' - ' + picker.endDate.format('YYYY-MM-DD'));
  });
  $('input[name="datefilter"]').on('cancel.daterangepicker', function(ev, picker) {
      $(this).val('');
  });
});
</script>
  <!-- <button type="submit" class="btn btn-primary">提交</button> -->
  <button id="button1" class="btn btn-default">提交</button>
        <label class="control-label">tips:(StartDay < DateRange <= EndDay)</label>
    </div>
  </div>
<!-- </form> -->
<br>
<!-- ：{{dateformat .Starttime "2006-01-02"}}-{{dateformat .Endtime "2006-01-02"}} -->

<h3>已完成</h3>
<div id="toolbar" class="btn-group">
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-plus"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-heart"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
    </button>
</div>
<table id="table"
        data-toggle="table"
       data-url="/completed"
       data-search="true"
       data-show-refresh="true"
       data-show-toggle="true"
       data-show-columns="true"
       data-toolbar="#toolbar"
       data-query-params="queryParams"
       >
    <thead>        
    <tr>
        <th data-formatter="index1">#</th>
        <th data-field="ProjectNumber">项目编号</th>
        <th data-field="ProjectName" data-sortable="true">项目名称</th>
        <th data-field="DesignStage" data-sortable="true">项目阶段</th>
        <th data-field="Tnumber">成果编号</th>
        <th data-field="Name">成果名称</th>
        <th data-field="Category" data-sortable="true">成果类型</th>
        <th data-field="Page">成果计量单位</th>
        <th data-field="Count">成果数量</th>
        <th data-field="Drawn">编制、绘制</th>
        <th data-field="Designd">设计</th>
        <th data-field="Checked">校核</th>
        <th data-field="Examined">审查</th>
        <th data-field="Date" data-formatter="localDateFormatter">出版</th>
<!--         <th data-field="action" data-formatter="actionFormatter" data-events="actionEvents">操作</th> -->
      </tr>
    </thead>
</table>

<!-- <div class="form-group">   
      <input type="button" class="btn btn-primary" name="insert" value="在线添加" onclick="insertNewRow()"/>
          <form id="form1" class="form-inline" method="post" action="/import_xls_catalog" enctype="multipart/form-data">
            <div class="form-group">
              <label>选择成果登记数据文件(Excel)
              <input type="file" class="form-control" name="catalog" id="catalog"></label>
              <br/>
              </div>
            <button type="submit" class="btn btn-primary" onclick="return import_xls_catalog();">提交</button>
          </form>
</div> --> 

<h3>已提交</h3>
<div id="running" class="btn-group">
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-plus"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-heart"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
    </button>
</div>
<table id="table1"
        data-toggle="table"
       data-url="/running"
       data-search="true"
       data-show-refresh="true"
       data-show-toggle="true"
       data-show-columns="true"
       data-toolbar="#running"
       data-query-params="queryParams"
       data-sort-name="ProjectName"
       data-sort-order="desc"
       >
    <thead>        
    <tr>
        <th data-formatter="index1">#</th>
        <th data-field="ProjectNumber">项目编号</th>
        <th data-field="ProjectName" data-sortable="true">项目名称</th>
        <th data-field="DesignStage" data-sortable="true">项目阶段</th>
        <th data-field="Tnumber" data-sortable="true">成果编号</th>
        <th data-field="Name">成果名称</th>
        <th data-field="Category" data-sortable="true">成果类型</th>
        <th data-field="Page">成果计量单位</th>
        <th data-field="Count">成果数量</th>
        <th data-field="Drawn">编制、绘制</th>
        <th data-field="Designd">设计</th>
        <th data-field="Checked">校核</th>
        <th data-field="Examined">审查</th>
        <th data-field="Date" data-formatter="localDateFormatter">出版</th>
      </tr>
    </thead>
</table>
<script>
function index1(value,row,index){
  // alert( "Data Loaded: " + index );
            return index+1
          }
function localDateFormatter(value) {
  return moment(value, 'YYYY-MM-DD').format('L');
}
function queryParams(params) {
  var date=$("#datefilter").val();
  var secid=$("#secid").val();
  var level=$("#level").val();
  // alert( "Data Loaded: " + date );
        params.datefilter=date;
        params.secid=secid;//传secid给后台，点击用户名，显示对应成果
        params.level=level;
        return params;
    }
$(function () {
        $('#button1').click(function () {
            //已完成
            $('#table').bootstrapTable('refresh', {url:'/completed'});
            //已提交
            $('#table1').bootstrapTable('refresh', {url:'/running'});
            //待提交
            $('#table2').bootstrapTable('refresh', {url:'/myself'});
            //待我处理的设计
            $('#table3').bootstrapTable('refresh', {url:'/designd'});
            //待我处理的校核
            $('#table4').bootstrapTable('refresh', {url:'/checked'});
            //待我处理的审查
            $('#table5').bootstrapTable('refresh', {url:'/examined'});

        });
    }); 
</script>

<h3>我发起，待提交</h3>
<div id="send" class="btn-group">
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-plus"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-heart"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
    </button>
</div>
<table id="table2"
        data-toggle="table"
       data-url="/myself"
       data-search="true"
       data-show-refresh="true"
       data-show-toggle="true"
       data-show-columns="true"
       data-toolbar="#send"
       data-query-params="queryParams"
       >
    <thead>        
    <tr>
        <th data-formatter="index1">#</th>
        <th data-field="ProjectNumber">项目编号</th>
        <th data-field="ProjectName">项目名称</th>
        <th data-field="DesignStage">项目阶段</th>
        <th data-field="Tnumber">成果编号</th>
        <th data-field="Name">成果名称</th>
        <th data-field="Category">成果类型</th>
        <th data-field="Page">成果计量单位</th>
        <th data-field="Count">成果数量</th>
        <th data-field="Drawn">编制、绘制</th>
        <th data-field="Designd">设计</th>
        <th data-field="Checked">校核</th>
        <th data-field="Examined">审查</th>
        <th data-field="Date" data-formatter="localDateFormatter">出版</th>
      </tr>
    </thead>
</table>

<h3>待我处理设计</h3>
<div id="designd" class="btn-group">
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-plus"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-heart"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
    </button>
</div>
<table id="table3"
        data-toggle="table"
       data-url="/designd"
       data-search="true"
       data-show-refresh="true"
       data-show-toggle="true"
       data-show-columns="true"
       data-query-params="queryParams"
       data-toolbar="#designd"
       >
    <thead>        
    <tr>
        <th data-formatter="index1">#</th>
        <th data-field="ProjectNumber">项目编号</th>
        <th data-field="ProjectName">项目名称</th>
        <th data-field="DesignStage">项目阶段</th>
        <th data-field="Tnumber">成果编号</th>
        <th data-field="Name">成果名称</th>
        <th data-field="Category">成果类型</th>
        <th data-field="Page">成果计量单位</th>
        <th data-field="Count">成果数量</th>
        <th data-field="Drawn">编制、绘制</th>
        <th data-field="Designd">设计</th>
        <th data-field="Checked">校核</th>
        <th data-field="Examined">审查</th>
        <th data-field="Date" data-formatter="localDateFormatter">出版</th>
      </tr>
    </thead>
</table>

<h3>待我处理校核</h3>
<div id="checked" class="btn-group">
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-plus"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-heart"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
    </button>
</div>
<table id="table4"
        data-toggle="table"
       data-url="/checked"
       data-search="true"
       data-show-refresh="true"
       data-show-toggle="true"
       data-show-columns="true"
       data-query-params="queryParams"
       data-toolbar="#checked">
    <thead>        
    <tr>
        <th data-formatter="index1">#</th>
        <th data-field="ProjectNumber">项目编号</th>
        <th data-field="ProjectName">项目名称</th>
        <th data-field="DesignStage">项目阶段</th>
        <th data-field="Tnumber">成果编号</th>
        <th data-field="Name">成果名称</th>
        <th data-field="Category">成果类型</th>
        <th data-field="Page">成果计量单位</th>
        <th data-field="Count">成果数量</th>
        <th data-field="Drawn">编制、绘制</th>
        <th data-field="Designd">设计</th>
        <th data-field="Checked">校核</th>
        <th data-field="Examined">审查</th>
        <th data-field="Date" data-formatter="localDateFormatter">出版</th>
      </tr>
    </thead>
</table>

<h3>待我处理审查</h3>
<div id="examined" class="btn-group">
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-plus"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-heart"></i>
    </button>
    <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
    </button>
</div>
<table id="table5"
        data-toggle="table"
       data-url="/examined"
       data-search="true"
       data-show-refresh="true"
       data-show-toggle="true"
       data-show-columns="true"
       data-query-params="queryParams"
       data-toolbar="#examined">
    <thead>        
    <tr>
        <th data-formatter="index1">#</th>
        <th data-field="ProjectNumber">项目编号</th>
        <th data-field="ProjectName">项目名称</th>
        <th data-field="DesignStage">项目阶段</th>
        <th data-field="Tnumber">成果编号</th>
        <th data-field="Name">成果名称</th>
        <th data-field="Category">成果类型</th>
        <th data-field="Page">成果计量单位</th>
        <th data-field="Count">成果数量</th>
        <th data-field="Drawn">编制、绘制</th>
        <th data-field="Designd">设计</th>
        <th data-field="Checked">校核</th>
        <th data-field="Examined">审查</th>
        <th data-field="Date" data-formatter="localDateFormatter">出版</th>
      </tr>
    </thead>
</table>
  
<tr>    
       <td colspan="4"><input type="button" class="btn btn-primary" name="insert" value="处&nbsp;&nbsp;&nbsp;&nbsp;理" onclick="ModifyRow()"/></td>    
       </tr>
       <br/>
       <br/>
</div>

<script type="text/javascript">
function import_xls_catalog(){
  var form1 = window.document.getElementById("form1");//获取form1对象
  form1.submit();
  $.ajax({
                        success:function(data,status){//数据提交成功时返回数据
                        // alert("添加“"+data+"”成功！(status:"+status+".)");
                        window.location.reload();
                        }
                    });
    return true;  //这个return必须放最后，前面的值才能传到后台    
   }

function insertNewRow(){
        // document.getElementById("iframepage").src="/secofficeshow?secid="+data.Id+"&level="+data.Level;
        window.open('/secofficeshow?secid='+{{.Secid}}+'&level='+{{.Level}}+'&key=editor');
        }
 function ModifyRow(){
        // document.getElementById("iframepage").src="/secofficeshow?secid="+data.Id+"&level="+data.Level;
        window.open('/secofficeshow?secid='+{{.Secid}}+'&level='+{{.Level}}+'&key=modify');
        }       

function actionFormatter(value, row, index) {
    return [
        '<a class="like" href="javascript:void(0)" title="Like">',
        '<i class="glyphicon glyphicon-heart"></i>',
        '</a>',
        '<a class="edit ml10" href="javascript:void(0)" title="Edit">',
        '<i class="glyphicon glyphicon-edit"></i>',
        '</a>',
        '<a class="remove ml10" href="javascript:void(0)" title="Remove">',
        '<i class="glyphicon glyphicon-remove"></i>',
        '</a>'
    ].join('');
}

window.actionEvents = {
    'click .like': function (e, value, row, index) {
        alert('You click like icon, row: ' + JSON.stringify(row));
        console.log(value, row, index);
    },
    'click .edit': function (e, value, row, index) {
        alert('You click edit icon, row: ' + JSON.stringify(row));
        console.log(value, row, index);
    },
    'click .remove': function (e, value, row, index) {
        alert('You click remove icon, row: ' + JSON.stringify(row));
        console.log(value, row, index);
    }
};
</script>
</body>
</html>
