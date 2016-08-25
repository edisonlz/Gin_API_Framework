function line_basic_b(cate,title,_names,_datas,_value_suffixs,divid,is_right){
    var series=new Array();
    for(var i=0;i <_names.length;i++){
        var json_data={};
        json_data.name=_names[i];
        json_data.data=_datas[i];
        if(_value_suffixs[i] == 1){
            json_data.tooltip={valueSuffix:'%'};
        }

        series.push(json_data);
    }

    var _legend={};
    if (is_right == 1){
        _legend.layout='vertical';
        _legend.align='right';
        _legend.verticalAlign='middle';
    }

    _legend.borderWidth=1;

    $(function () {
        $('#'+divid).highcharts({
            title: {
                text: title,
                x: -20 //center
            },
            subtitle: {
                text: '',
                x: -20
            },
            xAxis: {
                categories: cate,
                labels:{
                    rotation:-45

                }
            },
            yAxis: {
                title: {
                    text: ''
                },
                plotLines: [{
                    value: 0,
                    width: 1,
                    color: '#808080'
                }],
                labels:{
                    style: {
                        color: '#4572A7'
                    }
                }
            },
            tooltip: {
                shared:true,
                valueSuffix: ''
            },
            legend: _legend,
            series: series
        });
    });
}

function column_drilldown(cats,x_name,h_name,datas,colours,x_divid,y_title){

    $(function () {

        var colors = colours,
            categories=cats,
            name=x_name,
            title=h_name,
            divid=x_divid,
            data=datas;

        function setChart(name, categories, data, color) {
            chart.xAxis[0].setCategories(categories, false);
            chart.series[0].remove(false);
            chart.addSeries({
                name: name,
                data: data,
                color: color || 'white'
            }, false);
            chart.redraw();
        }

        var chart = $('#'+divid).highcharts({
            chart: {
                type: 'column'
            },
            title: {
                text: title
            },
            subtitle: {
                text: ''
            },
            xAxis: {
                categories: categories,
                labels:{
                    rotation:-45

                }
            },
            yAxis: {
                title: {
                    text: y_title
                }
            },
            plotOptions: {
                column: {
                    cursor: 'pointer',
                    point: {
                        events: {
                            click: function() {
                                var drilldown = this.drilldown;
                                if (drilldown) { // drill down
                                    setChart(drilldown.name, drilldown.categories, drilldown.data, drilldown.color);
                                } else { // restore
                                    setChart(name, categories, data);
                                }
                            }
                        }
                    },
                    dataLabels: {
                        enabled: true,
                        color: colors[0],
                        style: {
                            fontWeight: 'bold'
                        },
                        formatter: function() {
                            return this.y;
                        }
                    }
                }
            },
            tooltip: {
                formatter: function() {
                    var point = this.point,
                        s = this.x +': <b>'+ this.y +'</b><br/>';
                    return s;
                }
            },
            series: [{
                name: name,
                data: data,
                color: 'white'
            }],
            exporting: {
                enabled: false
            }
        })
            .highcharts(); // return chart
    });

}

function stacked_area(cate,title,series,divid) {
    $(function () {
        $('#'+divid).highcharts({
            chart: {
                type: 'area'
            },
            title: {
                text: title
            },
            xAxis: {
                categories: cate,
                tickmarkPlacement: 'on',
                title: {
                    enabled: false
                },
                labels:{
                    rotation:-45

                }
            },
            yAxis: {
                title: {
                    text: ''
                }
            },
            tooltip: {
                shared: true,
                valueSuffix: ''
            },
            plotOptions: {
                area: {
                    stacking: 'normal',
                    lineColor: '#666666',
                    lineWidth: 1,
                    marker: {
                        lineWidth: 1,
                        lineColor: '#666666'
                    }
                }
            },
            series:series
        });
    });
}

function combo_dual_axes(cate,title,_names,_datas,divid){
//    var y_colos = Highcharts.getOptions().colors;
    var y_colos=['#2f7ed8','#c42525','#8bbc21','#0d233a','#1aadce','#492970','#f28f43','#77a1e5','#a6c96a','#910000'];
    var series=new Array();
    var len=_names.length
    for(var i=0;i <len;i++){
        var json_data={};
        json_data.name=_names[i];
        json_data.data=_datas[i];
        json_data.color=y_colos[i];
        if (i!=len-1){
            json_data.type='area';
            json_data.yAxis=1;
        } else{
            json_data.type='spline';
            json_data.tooltip={valueSuffix:'%'};
        }
        series.push(json_data);
    }

    $(function () {
        $('#'+divid).highcharts({
            chart: {
                zoomType: 'xy'
            },
            title: {
                text: title
            },
            subtitle: {
                text: ''
            },
            xAxis: [{
                categories: cate,
                labels:{
                    rotation:-45

                }
            }],
            yAxis: [{ // Primary yAxis
                labels: {
                    format: '{value}%',
                    style: {
                        color: '#89A54E'
                    }
                },
                opposite:true,
                title: {
                    text: '',
                    style: {
                        color: '#89A54E'
                    }
                }
            }, { // Secondary yAxis
                title: {
                    text: '',
                    style: {
                        color: '#4572A7'
                    }
                },
                labels: {
                    format: '{value} ',
                    style: {
                        color: '#4572A7'
                    }
                }
            }],
            tooltip: {
                shared: true
            },
            series:series
        });
    });
}


function combo_dual_axes_push(cate,title,_names,_datas,divid){
    var y_colos = Highcharts.getOptions().colors;
//    var y_colos=['#4572A7','#cc0000','#89A54E']
    var series=new Array();
    var len=_names.length
    for(var i=0;i <len;i++){
        var json_data={};
        json_data.name=_names[i];
        json_data.data=_datas[i];
        json_data.color=y_colos[i];
        if (i <= 4){
            json_data.type='area';
            json_data.yAxis=1;
        } else{
            json_data.type='spline';
            json_data.tooltip={valueSuffix:'%'};
        }
        series.push(json_data);
    }

    $(function () {
        $('#'+divid).highcharts({
            chart: {
                zoomType: 'xy'
            },
            title: {
                text: title
            },
            subtitle: {
                text: ''
            },
            xAxis: [{
                categories: cate,
                labels:{
                    rotation:-45

                }
            }],
            yAxis: [{ // Primary yAxis
                labels: {
                    format: '{value}%',
                    style: {
                        color: '#89A54E'
                    }
                },
                opposite:true,
                title: {
                    text: '',
                    style: {
                        color: '#89A54E'
                    }
                }
            }, { // Secondary yAxis
                title: {
                    text: '',
                    style: {
                        color: '#4572A7'
                    }
                },
                labels: {
                    format: '{value} ',
                    style: {
                        color: '#4572A7'
                    }
                }
            }],
            tooltip: {
                shared: true
            },
            legend: {
                layout: 'vertical',
                align: 'right',
                verticalAlign: 'middle',
                borderWidth: 1
            },
            series:series
        });
    });
}

function checkAll(e, name) {
    var items = document.getElementsByName(name);
    for(var i=0;i<items.length;i++) {
        items[i].checked = e.checked;
    }
}

function line_basic_time(cate,title,_names,_datas,divid,is_right){
    var series=new Array();
    for(var i=0;i <_names.length;i++){
        var json_data={};
        json_data.name=_names[i];
        json_data.data=_datas[i];
        series.push(json_data);
    }

    var _legend={};
    if (is_right == 1){
        _legend.layout='vertical';
        _legend.align='right';
        _legend.verticalAlign='middle';
    }

    _legend.borderWidth=1;


    $(function () {
        $('#'+divid).highcharts({
            title: {
                text: title,
                x: -20 //center
            },
            subtitle: {
                text: '',
                x: -20
            },
            xAxis: {
                categories: cate
//                labels:{
//                    rotation:-45
//
//                }
            },
            yAxis: {
                title: {
                    text: ''
                },
                plotLines: [{
                    value: 0,
                    width: 1,
                    color: '#808080'
                }],
                labels:{
                    style: {
                        color: '#4572A7'
                    }
                }
            },
            tooltip: {
                valueSuffix: '',
                shared:true
            },
//            legend: {
//                borderWidth: 1,
//                layout: 'vertical',
//                align: 'right',
//                verticalAlign: 'middle'
//
//            },
            series: series
        });
    });
}


function parseObj( strData ){
    return (new Function( "return " + strData ))();
}

//小图缩放到最小，对应到统计首页的图函数
function littlelinebox(_id,_data){

    _data = parseObj(_data);

    $(function () {

        $('#'+_id).highcharts({

            chart: {

                //margin: [0, 0, 0, 0],

                //borderColor:'#F4F4F4',

                backgroundColor: '#F4F4F4',

                type: 'line'



            },

            title: {

                text: '',

                x: -20 //center

            },

            subtitle: {

                text: '',

                x: -20

            },

            xAxis: {

                tickWidth: 0,

                lineColor: '#F4F4F4',

                labels: {

                    enabled: false

                },

            },

            yAxis: {

                lineColor: '#F4F4F4',

                gridLineColor: '#F4F4F4',

                labels: {

                    enabled: false

                },

                title: {

                    text: ''

                },

                plotLines: [{

                    value: 0,

                    width: 0,

                    color: '#F4F4F4'

                }]

            },

            tooltip: {

                enabled:false,

            },

            legend: {

                floating: true,

                enabled:false,

                layout: 'vertical',

                align: 'right',

                verticalAlign: 'top',

                x: -10,

                y: 100,

                borderWidth: 0

            },plotOptions: {

                series: {

                    states: {

                        hover: {

                            enabled: false

                        }

                    }

                }

            },

            series: [{

                name: '',

                data: _data,

                marker: {

                    symbol: 'url(http://10.103.13.15:8000/static/dot.png)',

                    states: {

                        hover: {

                            enabled: false

                        },select: {

                            enabled: false

                        }

                    }

                }

            }]

        });

        //$("#highcharts-container").css("border","0px");

    });

}

//饼图函数
function PieChart(_id,_title,_name,_data){

    $(function () {
        $('#'+_id).highcharts({
            chart: {
                plotBackgroundColor: null,
                plotBorderWidth: null,
                plotShadow: false
            },
            title: {
                text: _title
            },
            tooltip: {
                pointFormat: '{series.name}: <b>{point.percentage:.1f}%</b>'
            },
            plotOptions: {
                pie: {
                    allowPointSelect: true,
                    cursor: 'pointer',
                    dataLabels: {
                        enabled: true,
                        color: '#000000',
                        connectorColor: '#000000',
                        format: '<b>{point.name}</b>: {point.percentage:.1f} %'
                    }
                }
            },
            series: [{
                type: 'pie',
                name: _name,
                data: _data
            }]
        });
    });
}

//Highcharts图中多跟柱装图函数
function Column_basic(cate,title,_names,_datas,divid){
    var series=new Array();
    var y_colos=['#c42525','#2f7ed8','#8bbc21','#0d233a','#1aadce','#492970','#f28f43','#77a1e5','#a6c96a','#910000'];
    for(var i=0;i <_names.length;i++){
        var json_data={};
        json_data.name=_names[i];
        json_data.data=_datas[i];
        json_data.color=y_colos[i]
        series.push(json_data);
    }

    $(function () {
        $('#'+divid).highcharts({
            chart: {
                type: 'column'
            },
            title: {
                text: title
            },
            subtitle: {
                text: ''
            },
            xAxis: {
                categories: cate,
                labels: {
                    rotation: -45,
                    align: 'right',
                    style: {
                        fontSize: '13px',
                        fontFamily: 'Verdana, sans-serif'
                    }
                }
            },
            yAxis: {
                min: 0,
                title: {
                    text: ''
                }
            },
            tooltip: {
                headerFormat: '<span style="font-size:10px">{point.key}</span><table>',
                pointFormat: '<tr><td style="color:{series.color};padding:0">{series.name}: </td>' +
                    '<td style="padding:0"><b>{point.y:.0f}</b></td></tr>',
                footerFormat: '</table>',
                shared: true,
                useHTML: true
            },
            plotOptions: {
                column: {
                    pointPadding: 0,
                    borderWidth: 0
                }
            },
            series: series
        });
    });

}

//Highcharts柱状图 只有一个柱子的图
function Column_rotated_labels(_id,_title,_cat,_data,_tip){
//    var y_colos=['#2f7ed8','#c42525','#8bbc21','#0d233a','#1aadce','#492970','#f28f43','#77a1e5','#a6c96a','#910000'];

    $(function () {
        $('#'+_id).highcharts({
            chart: {
                type: 'column',
                margin: [ 50, 50, 100, 80]
            },
            title: {
                text: _title
            },
            xAxis: {
                categories:_cat,
                labels: {
                    rotation: -45,
                    align: 'right',
                    style: {
                        fontSize: '13px',
                        fontFamily: 'Verdana, sans-serif'
                    }
                }
            },
            yAxis: {
                min: 0,
                title: {
                    text: ''
                }
            },
            legend: {
                enabled: false
            },
            tooltip: {
                pointFormat: _tip+':<b>{point.y:.1f}</b>'
            },
            series: [{
                name: 'Population',
                data: _data,
                color:'#2f7ed8',
                dataLabels: {
                    enabled: false,
                    rotation: -90,
                    color: '#FFFFFF',
                    align: 'right',
                    x: 4,
                    y: 10
//                    style: {
//                        fontSize: '13px',
//                        fontFamily: 'Verdana, sans-serif',
//                        textShadow: '0 0 3px black'
//                    }
                }
            }]
        });
    });
}


//为自动联想词代码专门写的函数,复用很高
function magicmenu(divid){

    (function( $ ) {
        $.widget( "custom.combobox", {
            _create: function() {
                this.wrapper = $( "<span>" )
                    .addClass( "custom-combobox" )
                    .insertAfter( this.element );

                this.element.hide();
                this._createAutocomplete();
                this._createShowAllButton();
            },

            _createAutocomplete: function() {
                var selected = this.element.children( ":selected" ),
                    value = selected.val() ? selected.text() : "";

                this.input = $( "<input onblur=\"get_game_version_list()\">" )
                    .appendTo( this.wrapper )
                    .val( value )
                    .attr( "title", "" )
                    .addClass( "custom-combobox-input ui-widget ui-widget-content ui-state-default ui-corner-left" )
                    .autocomplete({
                        delay: 0,
                        minLength: 0,
                        source: $.proxy( this, "_source" )
                    })
                    .tooltip({
                        tooltipClass: "ui-state-highlight"
                    });

                this._on( this.input, {
                    autocompleteselect: function( event, ui ) {
                        ui.item.option.selected = true;
                        this._trigger( "select", event, {
                            item: ui.item.option
                        });
                    },

                    autocompletechange: "_removeIfInvalid"
                });
            },

            _createShowAllButton: function() {
                var input = this.input,
                    wasOpen = false;

                $( "<a>" )
                    .attr( "tabIndex", -1 )
                    .attr( "title", "Show All Items" )
                    .tooltip()
                    .appendTo( this.wrapper )
                    .button({
                        icons: {
                            primary: "ui-icon-triangle-1-s"
                        },
                        text: false
                    })
                    .removeClass( "ui-corner-all" )
                    .addClass( "custom-combobox-toggle ui-corner-right" )
                    .mousedown(function() {
                        wasOpen = input.autocomplete( "widget" ).is( ":visible" );
                    })
                    .click(function() {
                        input.focus();

                        // Close if already visible
                        if ( wasOpen ) {
                            return;
                        }

                        // Pass empty string as value to search for, displaying all results
                        input.autocomplete( "search", "" );
                    });
            },

            _source: function( request, response ) {
                var matcher = new RegExp( $.ui.autocomplete.escapeRegex(request.term), "i" );
                response( this.element.children( "option" ).map(function() {
                    var text = $( this ).text();
                    if ( this.value && ( !request.term || matcher.test(text) ) )
                        return {
                            label: text,
                            value: text,
                            option: this
                        };
                }) );
            },

            _removeIfInvalid: function( event, ui ) {

                // Selected an item, nothing to do
                if ( ui.item ) {
                    return;
                }

                // Search for a match (case-insensitive)
                var value = this.input.val(),
                    valueLowerCase = value.toLowerCase(),
                    valid = false;
                this.element.children( "option" ).each(function() {
                    if ( $( this ).text().toLowerCase() === valueLowerCase ) {
                        this.selected = valid = true;
                        return false;
                    }
                });

                // Found a match, nothing to do
                if ( valid ) {
                    return;
                }

                // Remove invalid value
                this.input
                    .val( "" )
                    .attr( "title", value + " didn't match any item" )
                    .tooltip( "open" );
                this.element.val( "" );
                this._delay(function() {
                    this.input.tooltip( "close" ).attr( "title", "" );
                }, 2500 );
                this.input.data( "ui-autocomplete" ).term = "";
            },

            _destroy: function() {
                this.wrapper.remove();
                this.element.show();
            }
        });
    })( jQuery );

    $(function() {
        $( "#"+divid ).combobox();
        $( "#toggle" ).click(function() {
            $( "#"+divid ).toggle();
        });
    });

}

//为适配各个操作系统函数
function check_os() {

    var os_type = "MS Windows";

    windows = (navigator.userAgent.indexOf("Windows",0) != -1)?1:0;

    mac = (navigator.userAgent.indexOf("Mac",0) != -1)?1:0;

    linux = (navigator.userAgent.indexOf("Linux",0) != -1)?1:0;

    unix = (navigator.userAgent.indexOf("X11",0) != -1)?1:0;



    if (windows) os_type = "MS Windows";

    else if (mac) os_type = "Apple mac";

    else if (linux) os_type = "Lunix";

    else if (unix) os_type = "Unix";



    return os_type;

}

//为下载页面表格到Excel新写的函数
function getCSVData(tableID){

    var tabc = $('#'+tableID).html();

    tabc = tabc.replace(/<tr type[\s\S]*?<\/tr>/gi,'');

    $('#'+tableID).html(tabc);

    var os = check_os();

    if (os == "MS Windows" || os == "Lunix" || os == "Unix"){
        var csv_value=$('#'+tableID).table2CSV({delivery:'value'});
    }
    else if (os == "Apple mac"){
        var csv_value=$('#'+tableID).table2CSV({delivery:'value', separator : ','});
    }

    $("#csv_text").val(csv_value);
    document.formCSV.submit();

}


function combo_dual_axes_rerecommended(cate,title,_names,_datas,_value_suffixs,_image_type,divid,dtype){
//    var y_colos = Highcharts.getOptions().colors;
    var y_colos=['#2f7ed8','#c42525','#8bbc21','#0d233a','#1aadce','#492970','#f28f43','#77a1e5','#a6c96a','#910000'];
    var series=new Array();
    var len=_names.length
    for(var i=0;i <len;i++){
        var json_data={};
        json_data.name=_names[i];
        json_data.data=_datas[i];
        json_data.color=y_colos[i];

        if (dtype == 'channels'){
            if (_image_type[i] == 7 || _image_type[i] == 8 || _image_type[i] == 9) {
                json_data.type='spline';
            } else{
                json_data.type='area';
                json_data.yAxis=1;
            }
        }else{
            if (_image_type[i] == 3 || _image_type[i] == 4) {
                json_data.type='spline';
            } else{
                json_data.type='area';
                json_data.yAxis=1;
            }
        }

        if(_value_suffixs[i] == 1){
            json_data.tooltip={valueSuffix:'%'};
        }

        series.push(json_data);
    }

    $(function () {
        $('#'+divid).highcharts({
            chart: {
                zoomType: 'xy'
            },
            title: {
                text: title
            },
            subtitle: {
                text: ''
            },
            xAxis: [{
                categories: cate,
                labels:{
                    rotation:-45

                }
            }],
            yAxis: [{ // Primary yAxis
                labels: {
                    format: '{value}%',
                    style: {
                        color: '#89A54E'
                    }
                },
                opposite:true,
                title: {
                    text: '',
                    style: {
                        color: '#89A54E'
                    }
                }
            }, { // Secondary yAxis
                title: {
                    text: '',
                    style: {
                        color: '#4572A7'
                    }
                },
                labels: {
                    format: '{value} ',
                    style: {
                        color: '#4572A7'
                    }
                }
            }],
            tooltip: {
                shared: true
            },
            series:series
        });
    });
}


//Highcharts图中柱状堆叠图函数
function Stacted_column(cate,title,_names,_datas,divid){
    var series=new Array();
    var y_colos=['#B03060','#4682B4'];
    for(var i=0;i <_names.length;i++){
        var json_data={};
        json_data.name=_names[i];
        json_data.data=_datas[i];
        json_data.color=y_colos[i]
        series.push(json_data);
    }

    $(function () {
        Highcharts.setOptions({
        colors: [ '#B03060','#4682B4']
    });
        $('#'+divid).highcharts({
            chart: {
                type: 'column',
                borderWidth: 1,
                borderColor: '#DEDEDE'
            },
            title: {
                text: title
            },
            xAxis: {
                categories: cate,
                labels:{
                    rotation: -45,
                    align: 'right',
                    style:{
                        fontSize: '13px',
                        fontFamily: 'Verdana, scans-serif'
                    }
                }
            },
            yAxis: {
                min: 0,
                stackLabels: {
                    style: {
                        fontWeight: 'bold',
                        color: 'grey'
                    }
                },
                lineColor: '#DEDEDE',
                lineWidth: 1
            },
            legend: {
                layout: 'vertical',
                align: 'right',
                x: 0,
                verticalAlign: 'bottom',
                y: -70,
                backgroundColor: 'white',
                colors:['blue','red']
            },
//            tooltip: {
//                formatter: function() {
//                    return '<b>'+ this.x +'</b><br/>'+
//                        this.series.name +': '+ this.y +'<br/>'+
//                        'Total: '+ this.point.stackTotal;
//                }
//            },
            tooltip: {
                headerFormat: '<span style="font-size:10px">{point.key}</span><table>',
                pointFormat: '<tr><td style="color:{series.color};padding:0">{series.name}: </td>' +
                    '<td style="padding:0"><b>{point.y:.0f}</b></td></tr>',
                footerFormat: '</table>',
                shared: true,
                useHTML: true
            },
            plotOptions: {
                column: {
                    stacking: 'normal',
                    dataLabels: {
                        style: {
                            textShadow: '0 0 3px black, 0 0 3px black'
                        }
                    }
                }
            },
            series: series
        });
    });
}