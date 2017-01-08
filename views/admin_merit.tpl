<!-- iframe里定义成果类型和分值-->
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>Merit</title>
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script src="/static/js/bootstrap-treeview.js"></script>
  <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>

  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
  <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/>
  
  <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>

  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
  <script src="/static/js/tableExport.js"></script>
</head>
<body>

<script type="text/javascript">
  function index1(value,row,index){
      return index+1
  }

  // 改变点击行颜色
  $(function(){
     // $("#table").bootstrapTable('destroy').bootstrapTable({
     //     columns:columns,
     //     data:json
     // });
     $("#table0").on("click-row.bs.table",function(e,row,ele){
         $(".info").removeClass("info");
         $(ele).addClass("info");
         // rowid=row.Id;//全局变量
         // $("#details").show();
         // $('#table1').bootstrapTable('refresh', {url:'/admin/category/'+row.Id});
     });
     // $("#get").click(function(){
     //     alert("商品名称：" + getContent().TuanGouName);
     // })
  });
</script>

<div class="col-lg-12">
<h3>价值表</h3>
<div id="toolbar1" class="btn-group">
        <button type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
        </button>
        <button type="button" data-name="editorButton" id="editorButton" class="btn btn-default"> <i class="fa fa-edit">编辑</i>
        </button>
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
</div>

<table id="table0"
        data-toggle="table"
        data-url="/admin/merit"
        data-search="true"
        data-show-refresh="true"
        data-show-toggle="true"
        data-show-columns="true"
        data-toolbar="#toolbar1"
        data-query-params="queryParams"
        data-sort-name="ProjectName"
        data-sort-order="desc"
        data-page-size="5"
        data-page-list="[5, 25, 50, All]"
        data-unique-id="id"
        data-pagination="true"
        data-side-pagination="client"
        data-single-select="true"
        data-click-to-select="true"
        >
    <thead>        
      <tr>
        <!-- radiobox data-checkbox="true"-->
        <th data-width="10" data-radio="true"></th>
        <th data-formatter="index1">#</th>
        <th data-field="Title">价值名称</th>
        <th data-field="Mark">价值分值</th>
        <th data-field="List">价值选项</th>
        <th data-field="ListMark">选项分值</th>
        <!-- <th data-field="Iprole" data-title-tooltip="1-管理员;2-下载任意附件;3-下载pdf;4-查看成果">权限等级</th> -->
      </tr>
    </thead>
</table>

<script type="text/javascript">
        /*数据json*/
        var json =  [{"Id":"1","ProjCateName":"水利","ProjCateCode":"SL"},
                        {"Id":"2","ProjCateName":"电力","ProjCateCode":"DL"},
                        {"Id":"3","ProjCateName":"市政","ProjCateCode":"CJ"},
                        {"Id":"4","ProjCateName":"建筑","ProjCateCode":"JG"},
                        {"Id":"5","ProjCateName":"交通","ProjCateCode":"JT"},
                        {"Id":"6","ProjCateName":"境外","ProjCateCode":"JW"}];
        /*初始化table数据*/
        $(function(){
            $("#table0").bootstrapTable({
                data:json,
                // onClickRow: function (row, $element) {
                  // alert( "选择了行Id为: " + row.Id );
                  // rowid=row.Id//全局变量
                  // $('#table1').bootstrapTable('refresh', {url:'/admincategory?pid='+row.Id});
                // }
            });
        });

  $(document).ready(function() {
    $("#addButton").click(function() {
        $('#modalTable').modal({
        show:true,
        backdrop:'static'
        });
    })

    $("#editorButton").click(function() {
      var selectRow=$('#table0').bootstrapTable('getSelections');
      if (selectRow.length<1){
        alert("请先勾选！");
        return;
      }
      if (selectRow.length>1){
      alert("请不要勾选一个以上！");
      return;
      }
      $("input#cid").remove();
      var th1="<input id='cid' type='hidden' name='cid' value='" +selectRow[0].Id+"'/>"
      $(".modal-body").append(th1);//这里是否要换名字$("p").remove();
      $("#Title1").val(selectRow[0].Title);
      $("#Mark1").val(selectRow[0].Mark);
      $("#List1").val(selectRow[0].List);
      $("#ListMark1").val(selectRow[0].ListMark);

        $('#modalTable1').modal({
        show:true,
        backdrop:'static'
        });
    })

    $("#deleteButton").click(function() {
      var selectRow=$('#table0').bootstrapTable('getSelections');
      if (selectRow.length<=0) {
        alert("请先勾选！");
        return false;
      }
      var title=$.map(selectRow,function(row){
        return row.Title;
      })
      var ids="";
      for(var i=0;i<selectRow.length;i++){
        if(i==0){
          ids=selectRow[i].Id;
        }else{
          ids=ids+","+selectRow[i].Id;
        }  
      }
      $.ajax({
        type:"post",
        url:"/admin/merit/deletemerit",
        data: {ids:ids},
        success:function(data,status){
          alert("删除“"+data+"”成功！(status:"+status+".)");
          //删除已选数据
          $('#table0').bootstrapTable('remove',{
            field:'Title',
            values:title
          });
        }
      });  
    })
  })
</script>

<!-- 添加价值 -->
<div class="container">
  <form class="form-horizontal">
    <div class="modal fade" id="modalTable">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">添加价值</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">价值名称</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="Title"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">价值分值</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="Mark"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">价值选项</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="List"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">选项分值</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="ListMark"></div>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-primary" onclick="save()">保存</button>
          </div>  
        </div>
      </div>
    </div>
</form>
</div>
<script type="text/javascript">
  function save(){
    var Title  = $('#Title').val();
    var Mark  = $('#Mark').val();
    var List  = $('#List').val();
    var ListMark  = $('#ListMark').val();
    if (Title)
      {  
          $.ajax({
              type:"post",
              url:"/admin/merit/addmerit",
              data: {title:Title,mark:Mark,list:List,listmark:ListMark},
              success:function(data,status){
                alert("添加“"+data+"”成功！(status:"+status+".)");
               }
          });  
      }else{
        alert("名称不能为空");
      }  
        $('#modalTable').modal('hide');
        $('#table0').bootstrapTable('refresh', {url:'/admin/merit'});
  }

  function update(){
    var Title  = $('#Title1').val();
    var Mark  = $('#Mark1').val();
    var List  = $('#List1').val();
    var ListMark  = $('#ListMark1').val();
    var cid = $('#cid').val(); 
    if (Title)
      {  
          $.ajax({
              type:"post",
              url:"/admin/merit/updatemerit",
              data: {cid:cid,title:Title,mark:Mark,list:List,listmark:ListMark},
              success:function(data,status){
                alert("添加“"+data+"”成功！(status:"+status+".)");
               }
          });  
      }else{
        alert("名称不能为空");
      } 
        $('#modalTable1').modal('hide');
        $('#table0').bootstrapTable('refresh', {url:'/admin/merit'});
  }
</script>
<!-- 修改价值 -->
<div class="container">
  <form class="form-horizontal">
    <div class="modal fade" id="modalTable1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
            </button>
            <h3 class="modal-title">修改价值</h3>
          </div>
          <div class="modal-body">
            <div class="modal-body-content">
              <div class="form-group must">
                <label class="col-sm-3 control-label">价值名称</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="Title1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">价值分值</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="Mark1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">价值选项</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="List1"></div>
              </div>
              <div class="form-group must">
                <label class="col-sm-3 control-label">选项分值</label>
                <div class="col-sm-7">
                  <input type="text" class="form-control" id="ListMark1"></div>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
            <button type="button" class="btn btn-primary" onclick="update()">修改</button>
          </div>
        </div>
    </div>
  </div>
</form>
</div>


<script type="text/javascript">
 /*初始化table数据*/
 /*数据json*/
  var json1 = [{"Id":"1","ProjCateName":"规划","ProjCateCode":"A","ProjCateGrade":"1"},
              {"Id":"2","ProjCateName":"报告","ProjCateCode":"B","ProjCateGrade":"2"},
              {"Id":"3","ProjCateName":"图纸","ProjCateCode":"T","ProjCateGrade":"2"},
              {"Id":"4","ProjCateName":"水工","ProjCateCode":"5","ProjCateGrade":"3"},
              {"Id":"5","ProjCateName":"机电","ProjCateCode":"6","ProjCateGrade":"3"},
              {"Id":"6","ProjCateName":"施工","ProjCateCode":"7","ProjCateGrade":"3"}];
/*初始化table数据*/
// $(function(){
//     $("#table1").bootstrapTable({
//         data:json1
//     });
// });
        
// $('#editable td').on('change', function(evt, newValue) {
//     $.post( "script.php", { value: newValue })
//     .done(function( data ) {
//         alert( "Data Loaded: " + data );
//     });
// }); 
</script>

<br/>
</div>

</body>
</html>