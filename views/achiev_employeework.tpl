<!-- iframe里展示个人待处理的详细情况-->
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>待处理成果</title>
    <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
    <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
    <!-- <script src="/static/js/bootstrap-treeview.js"></script> -->
    <script type="text/javascript" src="/static/js/jquery.tablesorter.min.js"></script>
    <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/>
    <script type="text/javascript" src="/static/js/moment.min.js"></script>
    <script type="text/javascript" src="/static/js/daterangepicker.js"></script>
    <link rel="stylesheet" type="text/css" href="/static/css/daterangepicker.css"/>
    <script type="text/javascript" src="/static/bootstrap-datepicker/bootstrap-datepicker.js"></script>
    <script type="text/javascript" src="/static/bootstrap-datepicker/bootstrap-datepicker.zh-CN.js"></script>
    <link rel="stylesheet" type="text/css" href="/static/bootstrap-datepicker/bootstrap-datepicker3.css"/>

    <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-table.min.css"/>
    <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-editable.css"/>

    <script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
    <script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
    <script type="text/javascript" src="/static/js/bootstrap-table-editable.min.js"></script>
    <script type="text/javascript" src="/static/js/bootstrap-editable.js"></script>
    <script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>

    <link rel="stylesheet" type="text/css" href="/static/css/select2.css"/>
    <script type="text/javascript" src="/static/js/select2.js"></script>
    <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css"/>
    <script src="/static/js/tableExport.js"></script>
    <script src="/static/js/jquery.form.js"></script>

    <style>
      /*.form-group .datepicker{
        z-index: 9999;
      }*/
      i#delete
        {
          color:#C71585;
        }
    }
    </style>
  </head>

  <div class="col-lg-12">
    <h2>{{.UserNickname}}</h2>
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
            ranges : {
              'Today': [moment(), moment()],
              'Yesterday': [moment().subtract(1, 'days'), moment().subtract(1, 'days')],
              'Last 7 Days': [moment().subtract(6, 'days'), moment()],
              'Last 30 Days': [moment().subtract(29, 'days'), moment()],
              'This Month': [moment().startOf('month'), moment().endOf('month')],
              'Last Month': [moment().subtract(1, 'month').startOf('month'), moment().subtract(1, 'month').endOf('month')]
            },
            autoUpdateInput: false,
            locale: {
              cancelLabel: 'Clear'
            }
          });
          $('input[name="datefilter"]').on('apply.daterangepicker', function(ev,        picker){
            $(this).val(picker.startDate.format('YYYY-MM-DD') + ' - ' + picker.      endDate.   format('YYYY-MM-DD'));
          });
          $('input[name="datefilter"]').on('cancel.daterangepicker', function(ev,        picker)    {
            $(this).val('');
          });
        });
      </script>
      <button id="button" class="btn btn-default">提交</button>
      <label class="control-label">tips:(StartDay < DateRange <= EndDay)</label>
    </div>
    <br>

    <!-- <div class="form-inline">
      <input type='text' placeholder='项目编号' class="form-control" id='Pnumber' value='' size='4'/>
      <input type='text' placeholder='项目名称' class="form-control" id='Pname' value='' size='20'/>
      <select class="form-control" id='Stage'>
        <option>阶段：</option>
        <option>规划</option>
        <option>项目建议书</option>
        <option>可行性研究</option>
        <option>初步设计</option>
        <option>招标设计</option>
        <option>施工图</option>
      </select>
      <select class="form-control" id='Section'>
        <option>专业：</option>
        <option>水工</option>
        <option>施工</option>
        <option>预算</option>
      </select>
      <input type='text' placeholder='成果编号' class="form-control" id='Tnumber' value='' size='10'/>
      <input type='text' placeholder='成果名称' class="form-control" id='Name' value='' size='25'/>
    </div> -->
    <!-- <br/> -->
    <!-- <div class="form-inline">
      <select class="form-control" id='Category'>
        <option>成果类型：</option>
      </select>
      <input type='text' placeholder='数量' class="form-control" id='Count' value='' size='2'/>
      <input type='text' placeholder='绘制/编制' class="form-control" id="uname1" value='' list="cars1" size='7'/>
      <input type='text' placeholder='设计' class="form-control" id="uname2" value='' list="cars2" size='7'/>
      <input type='text' placeholder='校核' class="form-control" id="uname3" value='' list="cars3" size='7'/>
      <input type='text' placeholder='审查' class="form-control" id="uname4" value='' list="cars4" size='7'/>
      <input type='text' placeholder='绘制系数' class="form-control" id='Drawnratio' value='' size='4'/>
      <input type='text' placeholder='设计系数' class="form-control" id='Designdratio' value='' size='4'/>
      <input type='text' placeholder='出版日期' class='datepicker' id='Date' value='' size='7'/>
      <input type='button' class='btn btn-primary' name='update' value='添加' onclick='saveAddRow()'/>
      
    </div>
    <br/> -->

      <div id='datalistDiv'>
        <datalist id="cars1" name="cars1">
        </datalist>
      </div>
      <div id='datalistDiv'>
        <datalist id="cars2" name="cars2">
        </datalist>
      </div>
      <div id='datalistDiv'>
        <datalist id="cars3" name="cars3">
        </datalist>
      </div>
      <div id='datalistDiv'>
        <datalist id="cars4" name="cars4">
        </datalist>
      </div>
    <form id="form1" class="form-inline" method="post" action="/import_xls_catalog" enctype="multipart/form-data">
      <div class="form-group">
        <label>导入成果登记数据(Excel)
        <input type="file" class="form-control" name="catalog" id="catalog"></label>
        <br/>
      </div>
      <button type="submit" class="btn btn-primary" onclick="return import_xls_catalog();">提交</button>
    </form>
    <script type="text/javascript">
      function import_xls_catalog(){
          var file=$("#catalog").val();
          if(file!=""){  
              var form = $("form[id=form1]");
              var options  = {    
                  url:'import_xls_catalog',    
                  type:'post', 
                  success:function(data)    
                  {    
                    alert("导入数据："+data+"！")
                  }    
              };
             form.ajaxSubmit(options);
             return false;
          }else{
              alert("请选择文件！");
              return false; 
          }
      }
      function saveAddRow(){
        var newPnumber = $("#Pnumber").val();    
        var newPname = $("#Pname").val();    
        var newStage = $("#Stage option:selected").text();
        var newSection = $("#Section option:selected").text();
        var newTnumber = $("#Tnumber").val();
        var newName = $("#Name").val();
        var newCategory = $("#Category option:selected").text();
        
        var newCount = $("#Count").val();
        var newDrawn = $("#uname1").val();
        var newDesignd = $("#uname2").val();
        var newChecked = $("#uname3").val();
        var newExamined = $("#uname4").val();
        var newDrawnratio = $("#Drawnratio").val();
        var newDesigndratio = $("#Designdratio").val();
        var newDate = $("#Date").val();
        if(confirm("确定提交该行吗？")){    
          $.ajax({
          type:"post",//这里是否一定要用post？？？
          url:"/achievement/addcatalog",
          data: {Pnumber:newPnumber,Pname:newPname,Stage:newStage,Section:newSection,Tnumber:newTnumber,Name:newName,Category:newCategory,Count:newCount,Drawn:newDrawn,Designd:newDesignd,Checked:newChecked,Examined:newExamined,Drawnratio:newDrawnratio,Designdratio:newDesigndratio,Date:newDate},
            success:function(data,status){//数据提交成功时返回数据
              alert("添加“"+data+"”(status:"+status+".)");
              $('#table').bootstrapTable('refresh', {url:'/myself'});
            } 
          });
        }
      }     
      
      $(document).ready(function() {
        $("#addButton").click(function() {
          $('#modalTable').modal({
            show:true,
            backdrop:'static'
          });
        })
      })
    </script>
  <h3>我发起，待提交</h3>
    <div id="toolbar" class="btn-group">
        <button type="button" data-name="addButton" id="addButton" class="btn btn-default"> <i class="fa fa-plus">添加</i>
        </button>
        <button type="button" data-name="editorButton" id="editorButton" class="btn btn-default"> <i class="fa fa-edit">编辑</i>
        </button>
        <button type="button" data-name="deleteButton" id="deleteButton" class="btn btn-default">
        <i class="fa fa-trash">删除</i>
        </button>
    </div>
    <table id="table"
      data-query-params="queryParams"
      data-toolbar="#toolbar"
      data-search="true"
      data-show-refresh="true"
      data-show-toggle="true"
      data-show-columns="true"
      data-striped="true"
      data-clickToSelect="true"
      data-show-export="true"
      data-filter-control="true"
      >
    </table>
    <!-- 添加成果 -->
    <div class="container">
      <div class="form-horizontal">
        <div class="modal fade" id="modalTable">
          <div class="modal-dialog" style="width: 800px">
            <div class="modal-content">
              <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal">
                  <span aria-hidden="true">&times;</span>
                </button>
                <h3 class="modal-title">添加成果</h3>
              </div>
              <div class="modal-body">
                <div class="modal-body-content">
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">项目编号/名称</label>
                    <div class="col-sm-2">
                      <input type='text' placeholder='项目编号' class="form-control" id='Pnumber' value='' size='4'/>
                    </div> 
                  <!-- </div> -->
                  <!-- <div class="form-group must"> -->
                    <!-- <label class="col-sm-3 control-label">项目名称</label> -->
                    <div class="col-sm-6">
                      <input type='text' placeholder='项目名称' class="form-control" id='Pname' value='' size='20'/>
                    </div>    
                  </div>
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">项目阶段/专业</label>
                    <div class="col-sm-4">
                      <select class="form-control" id='Stage'>
                        <option>阶段：</option>
                        <option>规划</option>
                        <option>项目建议书</option>
                        <option>可行性研究</option>
                        <option>初步设计</option>
                        <option>招标设计</option>
                        <option>施工图</option>
                      </select>
                    </div>    
                  <!-- </div> -->
                  <!-- <div class="form-group must"> -->
                    <!-- <label class="col-sm-3 control-label">项目专业</label> -->
                    <div class="col-sm-4">
                      <select class="form-control" id='Section'>
                        <option>专业：</option>
                        <option>水工</option>
                        <option>施工</option>
                        <option>预算</option>
                      </select>
                    </div>    
                  </div>
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">成果编号/编号</label>
                    <div class="col-sm-3">
                      <input type='text' placeholder='成果编号' class="form-control" id='Tnumber' value='' size='10'/>
                    </div>    
                  <!-- </div> -->
                  <!-- <div class="form-group must"> -->
                    <!-- <label class="col-sm-3 control-label">成果名称</label> -->
                    <div class="col-sm-5">
                      <input type='text' placeholder='成果名称' class="form-control" id='Name' value='' size='25'/>
                    </div>    
                  </div>
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">成果类型/数量</label>
                    <div class="col-sm-4">
                      <select class="form-control" id='Category'>
                        <option>成果类型：</option>
                      </select>
                    </div>    
                  <!-- </div> -->
                  <!-- <div class="form-group must"> -->
                    <!-- <label class="col-sm-3 control-label">数量</label> -->
                    <div class="col-sm-4">
                      <input type='text' placeholder='数量' class="form-control" id='Count' value='' size='2'/>
                    </div>    
                  </div>
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">人员名称(拼音)</label>
                    <div class="col-sm-2">
                      <input type='text' placeholder='绘制/编制' class="form-control" id="uname1" value='' list="cars1" size='7'/>
                    </div>    
                  <!-- </div> -->
                  <!-- <div class="form-group must"> -->
                    <!-- <label class="col-sm-3 control-label">成果类型</label> -->
                    <div class="col-sm-2">
                      <input type='text' placeholder='设计' class="form-control" id="uname2" value='' list="cars2" size='7'/>
                    </div>    
                  <!-- </div> -->
                  <!-- <div class="form-group must"> -->
                    <!-- <label class="col-sm-3 control-label">成果类型</label> -->
                    <div class="col-sm-2">
                      <input type='text' placeholder='校核' class="form-control" id="uname3" value='' list="cars3" size='7'/>
                    </div>    
                  <!-- </div> -->
                  <!-- <div class="form-group must"> -->
                    <!-- <label class="col-sm-3 control-label">成果类型</label> -->
                    <div class="col-sm-2">
                      <input type='text' placeholder='审查' class="form-control" id="uname4" value='' list="cars4" size='7'/>
                    </div>    
                  </div>
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">绘制/设计系数</label>
                    <div class="col-sm-4">
                      <input type='text' placeholder='绘制/编制系数' class="form-control" id='Drawnratio' value='' size='4'/>
                    </div>    
                  <!-- </div> -->
                  <!-- <div class="form-group must"> -->
                    <!-- <label class="col-sm-3 control-label">成果类型</label> -->
                    <div class="col-sm-4">
                      <input type='text' placeholder='设计系数' class="form-control" id='Designdratio' value='' size='4'/>
                    </div>    
                  </div>
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">附件链接</label>
                    <div class="col-sm-8">
                      <input type='text' placeholder='http://' class="form-control" id='Link' value='http://' size='4'/>
                    </div>    
                  </div>
                  <div class="form-group must">
                    <label class="col-sm-3 control-label">成果说明</label>
                    <div class="col-sm-8">
                      <textarea class="form-control" rows="3" id='Content'></textarea>
                    </div>    
                  </div>
                  <div class="form-group">
                    <label class="col-sm-3 control-label">出版日期</label>
                    <div class="col-sm-3">
                      <span style="position: relative;z-index: 9999;">
                        <input type='text' placeholder='出版日期' class='datepicker' id='Date' value=''/>
                      </span>
                    </div>    
                  </div>
                </div>
              </div>
              <div class="modal-footer">
                <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
                <button type="button" class="btn btn-primary" onclick="saveAddRow()">保存</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <script type="text/javascript">
      $("#Date").datepicker({
        weekStart: 1,
        language: "zh-CN",
        autoclose: true,//选中之后自动隐藏日期选择框
        clearBtn: true,//清除按钮
        todayBtn: 'linked',//今日按钮
        format: "yyyy-mm-dd"//日期格式，详见 http://bootstrap-datepicker.readthedocs.org/en/release/options.html#format
      });
    </script>
    <h3>别人发起，我设计</h3>
    <div id="designd" class="btn-group">
        <button type="button" class="btn btn-default"> <i class="glyphicon    glyphicon-plus"></i>
        </button>
        <button type="button" class="btn btn-default"> <i class="glyphicon        glyphicon-heart"></i>
        </button>
        <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
        </button>
    </div>
    <table id="table1" 
      data-query-params="queryParams"
      data-toolbar="#designd"
      data-search="true"
      data-show-refresh="true"
      data-show-toggle="true"
      data-show-columns="true"
      data-striped="true"
      data-clickToSelect="true"
      data-show-export="true"
      data-filter-control="true"
       >
    </table>

<h3>别人发起，我校核</h3>
<div id="checked" class="btn-group">
        <button type="button" class="btn btn-default"> <i class="glyphicon glyphicon-plus"></i>
        </button>
        <button type="button" class="btn btn-default"> <i class="glyphicon glyphicon-heart"></i>
        </button>
        <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
        </button>
</div>
<table id="table2" 
      data-query-params="queryParams"
      data-toolbar="#checked"
      data-search="true"
      data-show-refresh="true"
      data-show-toggle="true"
      data-show-columns="true"
      data-striped="true"
      data-clickToSelect="true"
      data-show-export="true"
      data-filter-control="true"
       >
</table>
<h3>别人发起，我审查</h3>
<div id="examined" class="btn-group">
        <button type="button" class="btn btn-default"> <i class="glyphicon glyphicon-plus"></i>
        </button>
        <button type="button" class="btn btn-default"> <i class="glyphicon glyphicon-heart"></i>
        </button>
        <button type="button" class="btn btn-default">
        <i class="glyphicon glyphicon-trash"></i>
        </button>
</div>
<table id="table3" 
      data-query-params="queryParams"
      data-toolbar="#examined"
      data-search="true"
      data-show-refresh="true"
      data-show-toggle="true"
      data-show-columns="true"
      data-striped="true"
      data-clickToSelect="true"
      data-show-export="true"
      data-filter-control="true"
       >
</table>
<br/>
<br/>
</div>

<script type="text/javascript">
function actionFormatter(value, row, index) {
    return [
        '<a class="send" href="javascript:void(0)" title="提交">',
        '<i class="glyphicon glyphicon-step-forward"></i>',
        '</a>&nbsp;',
        '<a class="downsend" href="javascript:void(0)" title="退回">',
        '<i class="glyphicon glyphicon-step-backward"></i>',
        '</a>&nbsp;',
        '<a class="remove" href="javascript:void(0)" title="删除">',
        '<i id="delete" class="glyphicon glyphicon-remove"></i>',
        '</a>'
    ].join('');
}
// '<a class="edit ml10" href="javascript:void(0)" title="退回">','<i class="glyphicon glyphicon-edit"></i>','</a>'
window.actionEvents = {
    'click .send': function (e, value, row, index) {
        // alert('You click send icon, row: ' + JSON.stringify(row.Id));
        // alert(e);无值
        // alert(value);无值
        // alert(row);
        // alert(index);0~
        // console.log(value, row, index);
        if(confirm("确定提交该行吗？")){
          var removeline=$(this).parents("tr")
          //提交到后台进行修改数据库状态修改
            $.ajax({
            type:"post",//这里是否一定要用post？？？
            url:"/achievement/sendcatalog",
            data: {CatalogId:row.Id},
                success:function(data,status){//数据提交成功时返回数据
                removeline.remove();
                alert("提交“"+data+"”成功！(status:"+status+".)");
                }
            });  
        }
    },
    'click .downsend': function (e, value, row, index) {
        // alert('You click send icon, row: ' + JSON.stringify(row.Id));
        // alert(e);无值
        // alert(value);无值
        // alert(row);
        // alert(index);0~
        // console.log(value, row, index);
        if(confirm("确定退回该行吗？")){
        var removeline=$(this).parents("tr")
          //提交到后台进行修改数据库状态修改
            $.ajax({
            type:"post",//这里是否一定要用post？？？
            url:"/achievement/downsendcatalog",
            data: {CatalogId:row.Id},
                success:function(data,status){//数据提交成功时返回数据
                removeline.remove();
                alert("退回“"+data+"”成功！(status:"+status+".)");
                }
            });  
        }
    },

    // 'click .edit': function (e, value, row, index) {
    //     alert('You click edit icon, row: ' + JSON.stringify(row));
    //     console.log(value, row, index);
    // },
    'click .remove': function (e, value, row, index) {
        // alert('You click remove icon, row: ' + JSON.stringify(row));
        // console.log(value, row, index);
        if(confirm("确定删除该行吗？")){  
        var removeline=$(this).parents("tr")
        //提交到后台进行删除数据库
         // alert("欢迎您：" + name) 
            $.ajax({
            type:"post",//这里是否一定要用post？？？
            url:"/achievement/delete",
            data: {CatalogId:row.Id},
                success:function(data,status){//数据提交成功时返回数据
                removeline.remove();
                alert("删除“"+data+"”成功！(status:"+status+".)");
                }
            });  
        }
    }
};

//不提供删除功能的操作
function actionFormatter1(value, row, index) {
    return [
        '<a class="send" href="javascript:void(0)" title="提交">',
        '<i class="glyphicon glyphicon-step-forward"></i>',
        '</a>',
        '<a class="downsend" href="javascript:void(0)" title="退回">',
        '<i class="glyphicon glyphicon-step-backward"></i>',
        '</a>',
    ].join('');
}
//不提供删除功能的操作
window.actionEvents1 = {
    'click .send': function (e, value, row, index) {
        if(confirm("确定提交该行吗？")){
          var removeline=$(this).parents("tr")
          //提交到后台进行修改数据库状态修改
            $.ajax({
            type:"post",//这里是否一定要用post？？？
            url:"/achievement/sendcatalog",
            data: {CatalogId:row.Id},
                success:function(data,status){//数据提交成功时返回数据
                removeline.remove();
                alert("提交“"+data+"”成功！(status:"+status+".)");
                }
            });  
        }
    },
    'click .downsend': function (e, value, row, index) {
        if(confirm("确定退回该行吗？")){
        var removeline=$(this).parents("tr")
          //提交到后台进行修改数据库状态修改
            $.ajax({
            type:"post",//这里是否一定要用post？？？
            url:"/achievement/downsendcatalog",
            data: {CatalogId:row.Id},
                success:function(data,status){//数据提交成功时返回数据
                removeline.remove();
                alert("退回“"+data+"”成功！(status:"+status+".)");
                }
            });  
        }
    }
};
//这个是指定哪几个不能选的
function stateFormatter(value, row, index) {
    if (index === 2) {
        return {
            disabled: true
        };
    }
    if (index === 0) {
        return {
            disabled: true,
            checked: true
        }
    }
    return value;
}

//这个是导出的
// $(function () {
//   var $table = $('#table');
//   $('#toolbar').find('select').change(function () {
//     $table.bootstrapTable('refreshOptions', {
//       exportDataType: $(this).val()
//     });
//   });
// });
//这个是编辑表-2方法
// $(function () {
//     $('#table').bootstrapTable({
//         idField: 'ProjectNumber',
//         // pagination: true,
//         // search: true,
//         url: '/addinline',
//         columns: [{
//             field: 'Id',
//             title: '编号'
//         },
//         {
//             field: 'ProjectNumber',
//             title: '项目编号'
//         }, {
//             field: 'ProjectName',
//             title: '项目名称'
//         }],
//         onPostBody: function () {
//             $('#table').editableTableWidget({editor: $('<textarea>')});
//         }
//     });
// });
//在线编辑
// $(function () {
//   $('#table').bootstrapTable({
//     idField: 'ProjectNumber',
//     url: '/addinline',
//     columns: [{
//       field: 'Id',
//             title: '编号'
//         },
//         {
//       field: 'ProjectNumber',
//       title: 'ProjectNumber',
//       editable: {
//         type: 'text'
//       }
//     }, {
//       field: 'ProjectName',
//       title: 'ProjectName',
//       editable: {
//         type: 'address',
//         // var value={{.Ratio}}
//         display: function(value) {
//           if(!value) {
//             $(this).empty();
//             return; 
//           }
//           var html = '<b>' + $('<div>').text(value.Category).html() + '</b>, ' + $('<div>').text(value.Category).html() + ' st., bld. ' + $('<div>').text(value.Category).html();
//           $(this).html(html); 
//         }
//       }
//     }, {
//       field: 'description',
//       title: 'Description'
//     }]
//   });
// });
//待选择的修改*******不要删除
//我发起
$(function () {
    $('#table').bootstrapTable({
        idField: 'Id',
        url: '/achievement/myself',
        // striped: "true",
        columns: [
          {
            // field: 'Number',
            title: '序号',
            formatter:function(value,row,index){
            return index+1
          }
          },{
            field: 'ProjectNumber',
            title: '项目编号',
            sortable:'true',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter ProjectNumber' 
            }
          },{
            field: 'ProjectName',
            title: '项目名称',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter ProjectName'  
            }
          },{
            field: 'DesignStage',
            title: '阶段',
            editable: {
                type: 'select',
                source: ["规划", "项目建议书", "可行性研究", "初步设计", "招标设计", "施工图"],
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter DesignStage'  
            }
          },{
            field: 'Tnumber',
            title: '成果编号',
            sortable:'true',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter number'  
            }
          },{
            field: 'Name',
            title: '成果名称',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Name'  
            }
          },{
            field: 'Category',
            title: '成果类型',
            sortable:'true',
            editable: {
                type: 'select',
                source: {{.Select2}},//["$1", "$2", "$3"],
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Category' 
            }
          },{
            field: 'Count',
            title: '数量',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Count'  
            }
          },{
            field: 'Drawn',
            title: '制图/编制',
            editable: {
                type: 'select2', 
                source:{{.Userselect}},//'/regist/getuname1',
        // source: [
        //       {id: 'gb', text: 'Great Britain'},
        //       {id: 'us', text: 'United States'},
        //       {id: 'ru', text: 'Russia'}
        //    ],

        //'[{"id": "1", "text": "One"}, {"id": "2", "text": "Two"}]'

                select2: {
                  allowClear: true,
                  width: '150px',
                  placeholder: '请选择人名',
                  // multiple: true
                },//'/regist/getuname1',//这里用get方法，所以要换一个
                
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Drawn'  
            }
          },{
            field: 'Designd',
            title: '设计',
            editable: {
                type: 'select2', 
                source:{{.Userselect}},
                select2: {
                  allowClear: true,
                  width: '150px',
                  placeholder: '请选择人名',
                },
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Designd'  
            }
          },{
            field: 'Checked',
            title: '校核',
            editable: {
                type: 'select2', 
                source:{{.Userselect}},
                select2: {
                  allowClear: true,
                  width: '150px',
                  placeholder: '请选择人名',
                },
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Checked'  
            }
          },{
            field: 'Examined',
            title: '审查',
            editable: {
                type: 'select2', 
                source:{{.Userselect}},
                select2: {
                  allowClear: true,
                  width: '150px',
                  placeholder: '请选择人名',
                },
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Examined'  
            }
          },{
            field: 'Drawnratio',
            title: '制图比例',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Drawnratio'  
            }
          },{
            field: 'Designdratio',
            title: '设计比例',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Designdratio'  
            }
          },{
            field: 'Datestring',
            title: '出版(日/月/年)',
            // formatter:localDateFormatter,
            editable: {
                type: 'date',
                pk: 1,
                url: '/achievement/modifycatalog',
                // title: 'Enter ProjectNumber' 
                format: 'yyyy-mm-dd',    
                viewformat: 'dd/mm/yyyy',    
                datepicker: {
                    weekStart: 1,
                    todayBtn: 'linked'
                   }
                }
        },{
            field:'action',
            title: '操作',
            formatter:'actionFormatter',
            events:'actionEvents',
        }
        ]
    });
});
//我设计
$(function () {
    $('#table1').bootstrapTable({
        idField: 'Id',
        url: '/achievement/designd',
        // striped: "true",
        columns: [
          {
            // field: 'Number',
            title: '序号',
            formatter:function(value,row,index){
            return index+1
          }
          },{
            field: 'ProjectNumber',
            title: '项目编号',
            sortable:'true',
          },{
            field: 'ProjectName',
            title: '项目名称',
          },{
            field: 'DesignStage',
            title: '阶段',
          },{
            field: 'Tnumber',
            title: '成果编号',
            sortable:'true',
          },{
            field: 'Name',
            title: '成果名称',
          },{
            field: 'Category',
            title: '成果类型',
            sortable:'true',
          },{
            field: 'Count',
            title: '数量',
          },{
            field: 'Drawn',
            title: '制图/编制',
          },{
            field: 'Designd',
            title: '设计',
          },{
            field: 'Checked',
            title: '校核',
            editable: {
                type: 'select2', 
                source:{{.Userselect}},
                select2: {
                  allowClear: true,
                  width: '150px',
                  placeholder: '请选择人名',
                },
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Checked'  
            }
          },{
            field: 'Examined',
            title: '审查',
            editable: {
                type: 'select2', 
                source:{{.Userselect}},
                select2: {
                  allowClear: true,
                  width: '150px',
                  placeholder: '请选择人名',
                },
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Examined'  
            }
          },{
            field: 'Drawnratio',
            title: '制图比例',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Drawnratio'  
            }
          },{
            field: 'Designdratio',
            title: '设计比例',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Designdratio'  
            }
          },{
            field: 'Complex',
            title: '难度系数',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Complex'  
            }
          },{
            field: 'Datestring',
            title: '出版',
            // formatter:localDateFormatter,
            editable: {
                type: 'date',
                pk: 1,
                url: '/achievement/modifycatalog',
                // title: 'Enter ProjectNumber' 
                format: 'yyyy-mm-dd',    
                viewformat: 'dd/mm/yyyy',    
                datepicker: {
                    weekStart: 1,
                    todayBtn: 'linked'
                   }
                }
        },{
            field:'action',
            title: '操作',
            formatter:'actionFormatter1',
            events:'actionEvents1',
        }
        ]
    });
});

//我校核
$(function () {
    $('#table2').bootstrapTable({
        idField: 'Id',
        url: '/achievement/checked',
        // striped: "true",
        columns: [
          {
            // field: 'Number',
            title: '序号',
            formatter:function(value,row,index){
            return index+1
          }
          },{
            field: 'ProjectNumber',
            title: '项目编号',
            sortable:'true',
          },{
            field: 'ProjectName',
            title: '项目名称',
          },{
            field: 'DesignStage',
            title: '阶段',
          },{
            field: 'Tnumber',
            title: '成果编号',
            sortable:'true',
          },{
            field: 'Name',
            title: '成果名称',
          },{
            field: 'Category',
            title: '成果类型',
            sortable:'true',
          },{
            field: 'Count',
            title: '数量',
          },{
            field: 'Designd',
            title: '设计',
          },{
            field: 'Checked',
            title: '校核',
          },{
            field: 'Examined',
            title: '审查',
            editable: {
                type: 'select2', 
                source:{{.Userselect}},
                select2: {
                  allowClear: true,
                  width: '150px',
                  placeholder: '请选择人名',
                },
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Examined'  
            }
          },{
            field: 'Designdratio',
            title: '设计比例',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Designdratio'  
            }
          },{
            field: 'Checkedratio',
            title: '校核比例',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Checkedratio'  
            }
          },{
            field: 'Complex',
            title: '难度系数',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Complex'  
            }
          },{
            field: 'Datestring',
            title: '出版',
            // formatter:localDateFormatter,
            editable: {
                type: 'date',
                pk: 1,
                url: '/achievement/modifycatalog',
                // title: 'Enter ProjectNumber' 
                format: 'yyyy-mm-dd',    
                viewformat: 'dd/mm/yyyy',    
                datepicker: {
                    weekStart: 1,
                    todayBtn: 'linked'
                   }
                }
        },{
            field:'action',
            title: '操作',
            formatter:'actionFormatter1',
            events:'actionEvents1',
        }
        ]
    });
});

//我审查
$(function () {
    $('#table3').bootstrapTable({
        idField: 'Id',
        url: '/achievement/examined',
        // striped: "true",
        columns: [
          {
            // field: 'Number',
            title: '序号',
            formatter:function(value,row,index){
            return index+1
          }
          },{
            field: 'ProjectNumber',
            title: '项目编号',
            sortable:'true',
          },{
            field: 'ProjectName',
            title: '项目名称',
          },{
            field: 'DesignStage',
            title: '阶段',
          },{
            field: 'Tnumber',
            title: '成果编号',
            sortable:'true',
          },{
            field: 'Name',
            title: '成果名称',
          },{
            field: 'Category',
            title: '成果类型',
            sortable:'true',
          },{
            field: 'Count',
            title: '数量',
          },{
            field: 'Checked',
            title: '校核',
          },{
            field: 'Examined',
            title: '审查',
          },{
            field: 'Checkedratio',
            title: '校核比例',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Checkedratio'  
            }
          },{
            field: 'Examinedratio',
            title: '审查比例',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Examinedratio'  
            }
          },{
            field: 'Complex',
            title: '难度系数',
            editable: {
                type: 'text',
                pk: 1,
                url: '/achievement/modifycatalog',
                title: 'Enter Complex'  
            }
          },{
            field: 'Datestring',
            title: '出版',
            // formatter:localDateFormatter,
            editable: {
            type: 'date',
                pk: 1,
                url: '/achievement/modifycatalog',
                // title: 'Enter ProjectNumber' 
                format: 'yyyy-mm-dd',    
                viewformat: 'dd/mm/yyyy',    
                datepicker: {
                    weekStart: 1,
                    todayBtn: 'linked'
                }
            }
        },{
            field:'action',
            title: '操作',
            formatter:'actionFormatter1',
            events:'actionEvents1',
        }
        ]
    });
});

// var date={{.Starttime}};
// function list(value, row, index) {
 // return '<i class="glyphicon ' + icon + '"></i> ' + value;
    // return "<select data-index='row'><option>成果类型：</option></select>";
 // }
function localDateFormatter(value) {
     return moment(value, 'YYYY-MM-DD').format('L');
  }
function nameFormatter(value) {
    return '<a href="https://github.com/wenzhixin/' + value + '">' + value + '</a>';
}
//这个是显示时间选择
function datepicker(value) {
$(".datepicker").datepicker({
               language: "zh-CN",
               autoclose: true,//选中之后自动隐藏日期选择框
               clearBtn: true,//清除按钮
               todayBtn: 'linked',//今日按钮
               format: "yyyy-mm-dd"//日期格式，详见 http://bootstrap-datepicker.readthedocs.org/en/release/options.html#format
            });
}

function queryParams(params) {
  // var newPage = $("#txtPage").val();
  var date=$("#datefilter").val();
  params.datefilter=date;//"2016-09-10 - 2016-09-15";
        // params.your_param1 = 1; // add param1
        // params.your_param2 = 2; // add param2
        // console.log(JSON.stringify(params));
        // {"limit":10,"offset":0,"order":"asc","your_param1":1,"your_param2":2}
        return params;
    }

    // var $table = $('#table'),
    // $button = $('#button');
    $(function () {
        $('#button').click(function () {
            $('#table').bootstrapTable('refresh', {url:'/myself'});
            $('#table1').bootstrapTable('refresh', {url:'/designd'});
            $('#table2').bootstrapTable('refresh', {url:'/checked'});
            $('#table3').bootstrapTable('refresh', {url:'/examined'});
        });
    });    
// $(function () {
 // $('#button').click(function () {
      // var newPage = $("#txtPage").val();
            // var date=$("#datefilter").val();
            // params.datefilter=date;
            // alert( "Date Loaded: " + newPage);
            // $table.bootstrapTable('refresh', {url:'/addinline2'});
            // return params;
    // }); 
// });
// function queryParams() {
//         var params = {};
//         $('#toolbar').find('input[name]').each(function () {
//             params[$(this).attr('name')] = $(this).val();
//         });
//         return params;
//     }

// function queryParams(params) {
//             return {
//                 pageSize: params.pageSize,
//                 pageIndex: params.pageNumber,
//                 UserName: $("#txtName").val(),
//                 Birthday: $("#txtBirthday").val(),
//                 Gender: $("#Gender").val(),
//                 Address: $("#txtAddress").val(),
//                 name: params.sortName,
//                 order: params.sortOrder
//             };
//         }        
// 使用jQuery.post()方法传修改的数据到后台，这实际上是小菜一碟。

// $('#editable td').on('change', function(evt, newValue) {
//     $.post( "script.php", { value: newValue })
//     .done(function( data ) {
//         alert( "Data Loaded: " + data );
//     });
// });

// <input id="uname" name="uname" type="text" value="" class="form-control" placeholder="Enter account" list="cars"></div>
//         <div id='datalistDiv'>
//           <datalist id="cars" name="cars">//           </datalist>
//         </div>
      $(document).ready(function(){
        $("#sel_Province").change(function(){
          $.ajax({
            url: '<%=basePath%>areaAjax/getCity.do',
            data: "procode="+$("#sel_Province").val(),
            type: 'get',
            dataType:'json',
            error: function(data)
            {
              alert("加载json 文件出错！");
            },
            success: function(data)
            {
              for (var one in data){
                var name = data[one].name;
                var code = data[one].code;
                $("#sel_City").append("<option value="+code+">"+name+"</option>");
              }
            },
          });
        });
      });

      $(document).ready(function(){
      $.each({{.Select2}},function(i,d){
         $("#Category").append('<option value="' + i + '">'+d+'</option>');
         });
      });

      $('#uname1').attr("autocomplete","off"); 
      $(document).ready(function(){
        $("#uname1").keyup(function(event){
          // alert(event.keyCode);
          var uname1=document.getElementById("uname1");
        // if (uname.value.length==0)
         if (event.keyCode != 38 && event.keyCode != 40 && uname1.value.length==2){
          $.ajax({
                      type:"post",//这里是否一定要用post？？？
                      url:"/regist/getuname",
                      data: { uname: $("#uname").val()},
                      dataType:'json',//dataType:JSON,这种是jquerylatest版本的表达方法。不支持新版jquery。
                success:function(data,status){
                  $(".option").remove();
                  $.each(data,function(i,d){
                      $("#cars1").append('<option class="option" value="' + data[i].Username + '">' + data[i].Nickname + '</option>');
                  });
                }
          });
        }
      });
      }); 

      $('#uname2').attr("autocomplete","off");

      $(document).ready(function(){
        $("#uname2").keyup(function(event){
          var uname2=document.getElementById("uname2");
          // alert(event.keyCode);
         if (event.keyCode != 38 && event.keyCode != 40 && uname2.value.length==2){
          $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/regist/getuname",
                data: { uname: $("#uname").val()},
                dataType:'json',//dataType:JSON,这种是jquerylatest版本的表达方法。不支持新版jquery。
                success:function(data,status){
                  $(".option").remove();
                  $.each(data,function(i,d){
                      $("#cars2").append('<option class="option" value="' + data[i].Username + '">' + data[i].Nickname + '</option>');
                  });
                }
      });
                // $("#uname2").keydown(function(){
                //   $("option").remove();
                // }); 
    }
 });
}); 
$('#uname3').attr("autocomplete","off"); 
$(document).ready(function(){
  $("#uname3").keyup(function(event){
    var uname3=document.getElementById("uname3");
    // alert(event.keyCode);
   if (event.keyCode != 38 && event.keyCode != 40 && uname3.value.length==2){
    $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/regist/getuname",
                data: { uname: $("#uname").val()},
                dataType:'json',//dataType:JSON,这种是jquerylatest版本的表达方法。不支持新版jquery。
                success:function(data,status){
                  $(".option").remove();
                  $.each(data,function(i,d){
                      $("#cars3").append('<option class="option" value="' + data[i].Username + '">' + data[i].Nickname + '</option>');
                  });
                }
      });
                // $("#uname3").keydown(function(){
                //   $("option").remove();
                // }); 
    }
 });
}); 
$('#uname4').attr("autocomplete","off"); 
$(document).ready(function(){
  $("#uname4").keyup(function(event){
    var uname4=document.getElementById("uname4");
    // alert(event.keyCode);
   if (event.keyCode != 38 && event.keyCode != 40 && uname4.value.length==2){
    $.ajax({
                type:"post",//这里是否一定要用post？？？
                url:"/regist/getuname",
                data: { uname: $("#uname").val()},
                dataType:'json',//dataType:JSON,这种是jquerylatest版本的表达方法。不支持新版jquery。
                success:function(data,status){
                  $(".option").remove();
                  $.each(data,function(i,d){
                      $("#cars4").append('<option class="option" value="' + data[i].Username + '">' + data[i].Nickname + '</option>');
                  });
                }
      });
    //             $("#uname4").keydown(function(){
    //               $("option").remove();
    //             }); 
    }
 });
}); 
</script>

</body>
</html>