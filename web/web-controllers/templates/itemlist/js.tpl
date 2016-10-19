<script src="/static/js/jquery.form.js"></script>
<script src="/static/toast_message/jquery.toastmessage.js"></script>

<script type="text/javascript">

    $(document).ready(function () {
        /* list table sortbale */
        $("#sortable").sortable({
            revert: true,
            start: function(event, ui) {
                ui.item.startPos = ui.item.index();
            },
            stop: function(event, ui) {
                console.log("Start position: " + ui.item.startPos);
                console.log("New position: " + ui.item.index());
                if(ui.item.startPos!=ui.item.index())
                    $("#save_position_btn").removeClass("disabled");
            }
        });
        /* save positoin */
        $("#save_position_btn").click(function(){

            var item_ids = collect_module_ids_with_order();
//            console.log(item_ids);
            if(item_ids==''){
                alert('没有要排序的内容');
                return false;
            }
//            var $btn = $(this).button('loading');
            $("#gift_ids").val(item_ids);

            $('#form_positoin').submit();

        });

        /* save position form */
        var options = {
            success:  function(responseText, statusText){

                setTimeout(function () {$("#save_position_btn").button('reset')}, 1000);

                if(responseText.status=="success"){
                    $().toastmessage('showSuccessToast', '操作成功');
                    setTimeout(function () {location.reload()}, 1000);
                } else {
                    $().toastmessage('showErrorToast', "操作失败");
                }
            }
        };
        $('#form_positoin').ajaxForm(options);


        $('[data-toggle="tooltip"]').tooltip();


    });

    function collect_module_ids_with_order(){
        children = $("#sortable").children();
        ret = ''
        for(i=0;i<children.length;i++){
            child = children[i];
            ret += $(child).attr('value')+',';
        }
        return ret.substring(0, ret.length-1);
    }

</script>
