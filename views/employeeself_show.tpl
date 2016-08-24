<!-- iframe里展示个人可编辑的详细情况-->
<!DOCTYPE html>
<html>
<head>
 <meta charset="UTF-8">
  <title>个人情况汇总</title>
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
  <table class="table table-striped" id="orderTable" name="orderTable">
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
        <th>出版</th>
      </tr>
    </thead>

    <tbody>
      {{range $k,$v :=.Catalogs}}
      <tr id="row{{.Id}}">
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
        <td>{{dateformat .Data "2006-01-02"}}</td>
        <td><input type='button' class='btn btn-default' name='delete' value='删除' onclick='deleteSelectedRow("row{{.Id}}")'/> 
        <input type='button' class='btn btn-default' name='update' value='修改' onclick='updateSelectedRow("row{{.Id}}")' /></td> 
      </tr>
      {{end}}
    </tbody>
  </table>
  <!-- <input type="hidden" id="CategoryId" name="CategoryId" value="{{.CategoryId}}"/> -->
     <tr>    
       <td colspan="4"><input type="button" class="btn btn-default" name="insert" value="增加目录行" onclick="insertNewRow()"/></td>    
       </tr>
</div>



<script type="text/javascript">
//*********这个是编辑表格
var flag = 0;  //标志位，标志第几行  
        /*    
         *添加一个新行    
         */    
        function insertNewRow(){    
            //获得表格有多少行    
            var rowLength = $("#orderTable tr").length;  
            //这里的rowId 就是row加上标志位组合的，为了方便起见所以分开好一点。  
            var rowId = "row" + flag;  
            //每次都往低flag+1的下标出添加tr，因为append是往标签内追加，所以用after
            //"<td>￥<input type='text' id='txtDrawn"+flag+"' value='' size='10'/></td>"  
            var insertStr = "<tr id="+rowId+">" 
                         +      "<td><input type='text' placeholder='序号' id='txtIndex"+flag+"' value='' size='10'/></td>" 
                         +      "<td><input type='text' placeholder='项目编号' id='txtPnumber"+flag+"' value='' size='10'/></td>"  
                         +      "<td><input type='text' placeholder='项目名称' id='txtPname"+flag+"' value='' size='10'/></td>"  
                         +      "<td><input type='text' placeholder='阶段' id='txtStage"+flag+"' value='' size='10'/></td>"
                         +      "<td><input type='text' placeholder='成果编号' id='txtTnumber"+flag+"' value='' size='10'/></td>" 
                         +      "<td><input type='text' placeholder='成果名称' id='txtName"+flag+"' value='' size='10'/></td>" 
                         +      "<td><input type='text' placeholder='成果类型' id='txtCategory"+flag+"' value='' size='10'/></td>" 
                         +      "<td><input type='text' placeholder='计量单位' id='txtPage"+flag+"' value='' size='10'/></td>" 
                         +      "<td><input type='text' placeholder='数量' id='txtCount"+flag+"' value='' size='10'/></td>" 
                         +      "<td><input type='text' placeholder='编制/绘制' id='txtDrawn"+flag+"' value='' size='10'/></td>" 
                         +      "<td><input type='text' placeholder='设计' id='txtDesignd"+flag+"' value='' size='10'/></td>"
                         +      "<td><input type='text' placeholder='校核' id='txtChecked"+flag+"' value='' size='10'/></td>"
                         +      "<td><input type='text' placeholder='审查' id='txtExamined"+flag+"' value='' size='10'/></td>"
                         +      "<td><input type='text' placeholder='出版' id='txtData"+flag+"' value='' size='10'/></td>"
                         +      "<td><input type='button' class='btn btn-default' name='delete' value='删除' onclick='deleteSelectedRow(\""+rowId+"\")'/> <input type='button' class='btn btn-default' name='update' value='确定' onclick='saveAddRow(\""+rowId+"\",\""+flag+"\")' /></td>"                   
                         + "</tr>";  
            $("#orderTable tr:eq("+(rowLength-1)+")").after(insertStr);  //这里之所以减2 ，是因为减去底部的一行和顶部一行，剩下的为开始插入的索引。  
            flag++;  
        }    
  
        /*    
         *删除选中的行    
         */    
         function deleteSelectedRow(rowId){    
            //根据rowId查询出该行所在的行索引    
            if(confirm("确定删除该行吗？")){    
                $("#"+rowId).remove();    //这里需要注意删除一行之后 我的标志位没有-1，因为如果减一，那么我再增加一行的话，可能会导致我的tr的id重复，不好维护。
                // 提交到后台进行删除数据库
                    // alert("欢迎您：" + name) 
                    $.ajax({
                    type:"post",//这里是否一定要用post？？？
                    url:"/achievement/delete",
                    data: {CatalogId:rowId},
                        success:function(data,status){//数据提交成功时返回数据
                        alert("删除“"+data+"”成功！(status:"+status+".)");
                        }
                    });  
            }       
         }    
          
         /*    
          *修改选中的行    
          */    
         function updateSelectedRow(rowId){
            var oldIndex = $("#"+rowId+" td:eq(0)").html();
            var oldPnumber = $("#"+rowId+" td:eq(1)").html();  
            var oldPname = $("#"+rowId+" td:eq(2)").html();  
            var oldStage = $("#"+rowId+" td:eq(3)").html();
            var oldTnumber = $("#"+rowId+" td:eq(4)").html();
            var oldName = $("#"+rowId+" td:eq(5)").html();
            var oldCategory = $("#"+rowId+" td:eq(6)").html();
            var oldPage = $("#"+rowId+" td:eq(7)").html();
            var oldCount = $("#"+rowId+" td:eq(8)").html();
            var oldDrawn = $("#"+rowId+" td:eq(9)").html();
            var oldDesignd = $("#"+rowId+" td:eq(10)").html();
            var oldChecked = $("#"+rowId+" td:eq(11)").html();
            var oldExamined = $("#"+rowId+" td:eq(12)").html();
            var oldData = $("#"+rowId+" td:eq(13)").html();
            // if(oldPrice != ""){//去掉第一个人民币符号  
            //     oldPrice = oldPrice.substring(1);  
            // }  
            var uploadStr = "<td><input type='text' id='txtIndex"+flag+"' value='"+oldIndex+"' size='10'/></td>"
                        + "<td><input type='text' id='txtPnumber"+flag+"' value='"+oldPnumber+"' size='10'/></td>"  
                        + "<td><input type='text' id='txtPname"+flag+"' value='"+oldPname+"' size='10'/></td>"  
                        + "<td><input type='text' id='txtStage"+flag+"' value='"+oldStage+"' size='10'/></td>"
                        + "<td><input type='text' id='txtTnumber"+flag+"' value='"+oldTnumber+"' size='10'/></td>"
                        + "<td><input type='text' id='txtName"+flag+"' value='"+oldName+"' size='10'/></td>"
                        + "<td><input type='text' id='txtCategory"+flag+"' value='"+oldCategory+"' size='10'/></td>"
                        + "<td><input type='text' id='txtPage"+flag+"' value='"+oldPage+"' size='10'/></td>"
                        + "<td><input type='text' id='txtCount"+flag+"' value='"+oldCount+"' size='10'/></td>"
                        + "<td><input type='text' id='txtDrawn"+flag+"' value='"+oldDrawn+"' size='10'/></td>"
                        + "<td><input type='text' id='txtDesignd"+flag+"' value='"+oldDesignd+"' size='10'/></td>"
                        + "<td><input type='text' id='txtChecked"+flag+"' value='"+oldChecked+"' size='10'/></td>"
                        + "<td><input type='text' id='txtExamined"+flag+"' value='"+oldExamined+"' size='10'/></td>"
                        + "<td><input type='text' id='txtData"+flag+"' value='"+oldData+"' size='10'/></td>"
                        + "<td><input type='button' class='btn btn-default' name='delete' value='删除' onclick='deleteSelectedRow(\""+rowId+"\")'/> <input type='button' class='btn btn-default' name='update' value='确定' onclick='saveUpdateRow(\""+rowId+"\",\""+flag+"\")' /></td>";  
            $("#"+rowId).html(uploadStr);  
         }    
  
         /*    
          *保存添加    
          */    
          function saveAddRow(rowId,flag){ 
            var newIndex = $("#txtIndex"+flag).val();
            var newPnumber = $("#txtPnumber"+flag).val();    
            var newPname = $("#txtPname"+flag).val();    
            var newStage = $("#txtStage"+flag).val();
            var newTnumber = $("#txtTnumber"+flag).val();
            var newName = $("#txtName"+flag).val();
            var newCategory = $("#txtCategory"+flag).val();
            var newPage = $("#txtPage"+flag).val();
            var newCount = $("#txtCount"+flag).val();
            var newDrawn = $("#txtDrawn"+flag).val();
            var newDesignd = $("#txtDesignd"+flag).val();
            var newChecked = $("#txtChecked"+flag).val();
            var newExamined = $("#txtExamined"+flag).val();
            var newData = $("#txtData"+flag).val();
            var saveStr = "<td>" + newIndex + "</td>"
                        + "<td>" + newPnumber + "</td>"  
                        + "<td>" + newPname + "</td>"  
                        + "<td>" + newStage + "</td>"
                        + "<td>" + newTnumber + "</td>"
                        + "<td>" + newName + "</td>"
                        + "<td>" + newCategory + "</td>"
                        + "<td>" + newPage + "</td>"
                        + "<td>" + newCount + "</td>"
                        + "<td>" + newDrawn + "</td>"
                        + "<td>" + newDesignd + "</td>"
                        + "<td>" + newChecked + "</td>"
                        + "<td>" + newExamined + "</td>"
                        + "<td>" + newData + "</td>"
                        + "<td><input type='button' class='btn btn-default' name='delete' value='删除' onclick='deleteSelectedRow(\""+rowId+"\")'/> <input type='button' class='btn btn-default' name='update' value='修改' onclick='updateSelectedRow(\""+rowId+"\")' /></td>";  
            $("#"+rowId).html(saveStr);//因为替换的时候只替换tr标签内的html 所以不用加上tr 
            // 这里再提交到后台保存起来update 
            if (newName)//如果返回的有内容  
                {  
                 // var pid = $('#CategoryId').val();
                    // alert("欢迎您：" + name) 
                    $.ajax({
                    type:"post",//这里是否一定要用post？？？
                    url:"/achievement/addcatalog",
                    data: {Pnumber:newPnumber,Pname:newPname,Stage:newStage,Tnumber:newTnumber,Name:newName,Category:newCategory,Page:newPage,Count:newCount,Drawn:newDrawn,Designd:newDesignd,Checked:newChecked,Examined:newExamined,Data:newData},
                        success:function(data,status){//数据提交成功时返回数据
                        alert("添加“"+data+"”成功！(status:"+status+".)");
                        }
                    });  
                }
          }

          /*    
          *保存修改    
          */
        function saveUpdateRow(rowId,flag){ 
            var newIndex = $("#txtIndex"+flag).val();
            var newPnumber = $("#txtPnumber"+flag).val();    
            var newPname = $("#txtPname"+flag).val();    
            var newStage = $("#txtStage"+flag).val();
            var newTnumber = $("#txtTnumber"+flag).val();
            var newName = $("#txtName"+flag).val();
            var newCategory = $("#txtCategory"+flag).val();
            var newPage = $("#txtPage"+flag).val();
            var newCount = $("#txtCount"+flag).val();
            var newDrawn = $("#txtDrawn"+flag).val();
            var newDesignd = $("#txtDesignd"+flag).val();
            var newChecked = $("#txtChecked"+flag).val();
            var newExamined = $("#txtExamined"+flag).val();
            var newData = $("#txtData"+flag).val();
            var saveStr = "<td>" + newIndex + "</td>"
                        + "<td>" + newPnumber + "</td>"  
                        + "<td>" + newPname + "</td>"  
                        + "<td>" + newStage + "</td>"
                        + "<td>" + newTnumber + "</td>"
                        + "<td>" + newName + "</td>"
                        + "<td>" + newCategory + "</td>"
                        + "<td>" + newPage + "</td>"
                        + "<td>" + newCount + "</td>"
                        + "<td>" + newDrawn + "</td>"
                        + "<td>" + newDesignd + "</td>"
                        + "<td>" + newChecked + "</td>"
                        + "<td>" + newExamined + "</td>"
                        + "<td>" + newData + "</td>"
                        + "<td><input type='button' class='btn btn-default' name='delete' value='删除' onclick='deleteSelectedRow(\""+rowId+"\")'/> <input type='button' class='btn btn-default' name='update' value='修改' onclick='updateSelectedRow(\""+rowId+"\")' /></td>";  
            $("#"+rowId).html(saveStr);//因为替换的时候只替换tr标签内的html 所以不用加上tr 
            // 这里再提交到后台保存起来update 
            if (newName)//如果返回的有内容  
                {  
                    $.ajax({
                    type:"post",//这里是否一定要用post？？？
                    url:"/achievement/modifycatalog",
                    data: {Pnumber:newPnumber,Pname:newPname,Stage:newStage,Tnumber:newTnumber,Name:newName,Category:newCategory,Page:newPage,Count:newCount,Drawn:newDrawn,Designd:newDesignd,Checked:newChecked,Examined:newExamined,Data:newData,CatalogId:rowId},
                        success:function(data,status){//数据提交成功时返回数据
                        alert("修改“"+data+"”成功！(status:"+status+".)");
                        }
                    });  
                }
          }

</script>




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
  $("table").tablesorter({sortList: [[13,0]]});
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