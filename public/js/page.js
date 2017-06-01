/**
 * User: Hiko
 * 微我 - 微我页js
 */
$(document).ready(function(){

    // 拉拽封面管理框
    var cover_manage_drag_even = function() {
        var cover_manage_container = $('#cover-manage-container');
        var cover_manage_header = $('#cover-manage-header');

        //管理框框宽高
        var box_width = cover_manage_container.width();
        var box_height = cover_manage_container.height();

        //浏览器宽高
        var win_width = $(window).width();
        var win_height = $(window).height();

        // 使管理框居中
        cover_manage_container.css({
            'left': (win_width - box_width) / 2
        });

        //菜单高度
        var menu_height = 80;

        //信息框左边到鼠标点下的左边水平距离、垂直距离； 管理框到浏览器左边水平距离、垂直距离；
        var initX = 0, initY = 0, offsetX = 0, offsetY = 0;

        //标记是否进行拉拽
        var dragging = false;

        cover_manage_header.mousedown(function(mouse_even) {
            dragging = true;
            initX = mouse_even.clientX - cover_manage_container[0].offsetLeft;
            initY = mouse_even.clientY - cover_manage_container[0].offsetTop + menu_height;
            this.setCapture && this.setCapture();

            //移动信息框事件 (因为该页面内有两处地方需要点击移动，所以在不同的情况下，给document的onmousemove附上不同的事件)
            document.onmousemove = function(mouse_even) {
                // console.log('offsetx: '+ offsetX + ', offsetY:' + offsetY);

                if(dragging) {
                    var mouse_even =  mouse_even || windown.even;

                    offsetX = mouse_even.clientX - initX;
                    offsetY = mouse_even.clientY - initY;

                    //判断拉拽位置限制
                    if(offsetX <= 0)
                    {
                        offsetX = 0;
                    }

                    if(offsetY <= 0)
                    {
                        offsetY = 0;
                    }

                    // 4 = 管理框border宽度(2px) * 2
                    if((offsetX + box_width) >= win_width + 4)
                    {
                        offsetX = win_width - box_width - 4;
                    }
                    if(offsetY + box_height >= win_height)
                    {
                        offsetY = win_height - box_height;
                    }

                    cover_manage_container.css({
                        "left":offsetX + "px",
                        "top":offsetY + "px"
                    });

                    return false;
                }
            };

            return false;
        });

        $('body').mouseup(function() {
            dragging = false;
        });


    };

    cover_manage_drag_even();

    //背景图片 Initialize Background Stretcher
    $(document).bgStretcher({
        //使用在html页面定义的全局变量：background_img - 用户的背景图片
        images: user_setting.background_img,
        imageWidth: 1024/*, imageHeight: 768*/

    });
    /*$('#toggleAnimation').click(function(){
         if($(this).val() == "Pause Animation"){
            $(this).val("Resume Animation");
            $(document).bgStretcher.pause();
         } else {
            $(this).val("Pause Animation");
            $(document).bgStretcher.play();
         }
     });*/

    /*if(navigator.userAgent.indexOf("MSIE 8.0")>0)
     {
        alert("ie8");
     }
     if(navigator.userAgent.indexOf("MSIE 9.0")>0)
     {
        alert("ie9");
     }*/


    //弹出个人中心..etc
    var top_menu_login_info = $('#top-menu-login-info');

    top_menu_login_info.mouseover(function() {
        $(this).find('.dropdown-menu').show();
    }).mouseout(function() {
        $(this).find('.dropdown-menu').hide();
    });

    /**
     * 管理封面图
     */
    var layout_cover_manage = $('#layout-cover-manage');
    var sys_support_box = $('#cover-manage-sys-support');
    var custom_box = $('#cover-manage-custom');
    var nav_sys_support = $('#cover-manage-nav-sys-support');
    var nav_custom = $('#cover-manage-nav-custom');

    $('#manage-cover-btn').click(function() {
        layout_cover_manage.fadeIn(100);
        var cover_nav = sys_support_pages.find('a');
        cover_nav.attr('class', 'btn');
        $(cover_nav[0]).attr('class', 'btn btn-inverse');

        // 1.1 初始化第一页 - 系统提供的封面图
        get_sys_covers();
    });

    $('#layout-close-cover-manage, #cover-manage-cancel').click(function() {

        // 还原原来的封面图
        $('.cover-item-selected').remove();
        var old_cover = old_custom_cover_input.val();
        cover_manage_custom_pic.attr('src', old_cover);
        var old_cover_height = parseInt(old_cover_height_input.val());

        bgStretcher.html('<ul><li><img width="100%" src="' + old_cover + '"></li></ul>');

        // 调整封面图高度
        change_cover_height(old_cover_height);

        // 调整滑块位置
        the_block.css({
            'left': old_cover_height - user_setting.default_cover_height
        });

        new_cover_id_input.val('');
        new_cover_input.val('');

        layout_cover_manage.fadeOut(100);
    });

    nav_sys_support.click(function() {
        custom_box.hide();
        sys_support_box.show();

        nav_custom.attr("class", "");
        $(this).attr("class", "active");
    });

    nav_custom.click(function() {
        sys_support_box.hide();
        custom_box.show();

        nav_sys_support.attr("class", "");
        $(this).attr("class", "active");

    });

    // 上传封面图片
    var cover_upload_btn = $('#cover-upload');
    var cover_upload_error_tip = $('#cover-select-error');
    var cover_px_tip = $('#cover-px-tip');
    var new_cover_id_input = $("#new-cover-id");
    var new_cover_input = $('#new-cover');
    var old_cover_height_input = $('#old-cover-height');
    var old_custom_cover_input =$('#old-custom-cover');
    var cancel_new_cover_btn = $('#cancel-new-upload-cover');
    var bgStretcher = $('#bgstretcher');
    var cover_manage_custom_pic = $('#cover-manage-custom-pic');

    cover_upload_btn.uploadify({
        /*'formData'     : {
            'timestamp' : '<?php echo $timestamp;?>',
            'token'     : '<?php echo md5('unique_salt' . $timestamp);?>'
        },*/
        'swf'      : 'public/css/uploadify.swf',
        'uploader' : user_setting.cover_up,
        'fileObjName': 'cover',
        'multi': true,
        'auto': true,
        'buttonText': '<span><i class="icon-picture icon-white"></i> 选择封面图</span>',
        'fileSizeLimit': 1024 * 5,
        'fileTypeDesc'       : '支持格式:jpg,png,jpeg.',
        'fileTypeExts'        : '*.jpg; *.png; *.jpeg',
        'overrideEvents' : [ 'onDialogClose', 'onUploadError', 'onSelectError' ],

        'onDialogClose': function(queueData) {
            if(queueData.filesQueued > 0) {
                cover_upload_btn.uploadify('disable', true);
            }

            cover_upload_error_tip.hide();
            cover_px_tip.show();
        },

        'onUploadStart' : function(file) {},

        'onCancel' : function(file) {
            cover_upload_btn.uploadify('disable', false);
        },

        'onUploadError': function(file, errorCode, errorMsg, errorString) {
            cover_upload_btn.uploadify('disable', false);

            $('.uploadify-queue-item').hide();
            cover_px_tip.hide();
            cover_upload_error_tip.html('<b>提示：</b>上传失败，请稍后再尝试。' + errorString).show();
        },

        'onQueueComplete':  function(queueData) {
            cover_upload_btn.uploadify('disable', false);
        },

        'onSelectError' : function(file, errorCode, errorMsg, errorString) {
            if(errorCode === -110) {
                $('.uploadify-queue-item').hide();
                cover_px_tip.hide();
                cover_upload_error_tip.show();
            }else {
                cover_upload_error_tip.html('<b>提示：</b>选择图片失败，请刷新网页之后再尝试。' + errorMsg).show();
            }

            return false;   
        },

        'onUploadSuccess': function (file, data, response) {
            if(!response) {
                cover_px_tip.hide();
                cover_upload_error_tip.show();
            }

            ret = eval("(" + data + ")");
            if(!ret.status) {
                cover_px_tip.hide();
                cover_upload_error_tip.html('<b>提示：</b>上传失败，请重新尝试。').show();
                return;
            }

            // 设置新界面项
            cover_manage_custom_pic.attr('src', ret.data.cover_src);
            new_cover_input.val(ret.data.cover_src);
            new_cover_id_input.val(ret.data.cover_id);
            cancel_new_cover_btn.show();
            bgStretcher.html('<ul><li><img width="100%" src="' + ret.data.cover_src + '"></li></ul>');
        }

    });

    // 撤销 - 刚刚上传的封面图 - 确认框
    var confirm_box = $('#weiwo-tip-box');
    var confirm_box_width = confirm_box.width();
    var confirm_box_height = confirm_box.height();
    var win_width = $(window).width();
    var win_height = $(window).height();

    cancel_new_cover_btn.click(function() {
        confirm_box.css({
            'left': (win_width - confirm_box_width) / 2 + 'px',
            'top': (win_height - confirm_box_height) / 2 + 'px'
        });

        confirm_box.show();
    });

    $('#do-cancel-new-cover').click(function() {
        var old_cover = old_custom_cover_input.val();
        // console.log(old_cover);
        cover_manage_custom_pic.attr('src', old_cover);
        bgStretcher.html('<ul><li><img width="100%" src="' + old_cover + '"></li></ul>');

        // 还原原来的封面图
        new_cover_id_input.val('');
        new_cover_input.val('');
        confirm_box.hide();
        cancel_new_cover_btn.hide();
    });

    $('.close-weiwo-tip-box').click(function() {
        confirm_box.hide();
    });

    /**
     * 封面图高度 - 拉拽
     * @type {*|jQuery|HTMLElement}
     */
    var bgContainer = $('#container');
    var bgContainerHeight = parseInt(bgContainer.height());
    var the_block = $('#custom-cover-height-block');
    var the_bar = $('#custom-cover-height-bar');

    var cover_height_drag = function() {

        var block_left = the_block.css('left');
        if(block_left == undefined) { // 表示该page不是当前用户的
            return;
        }
        block_left = block_left.substring(0, block_left.length - 2);
        var block_width = the_block.width();
        var bar_width = the_bar.width();
        var interval_left = 0;

        //标记是否进行拉拽
        var dragging = false;

        var mouse_down_left = 0;

        the_block.mousedown(function(mouse_even) {
            dragging = true;
            block_left = the_block.css('left');
            block_left = parseInt(block_left.substring(0, block_left.length - 2));
            mouse_down_left = mouse_even.clientX;

            this.setCapture && this.setCapture();

            //移动"块"事件 (因为该页面内有两处地方需要点击移动，所以在不同的情况下，给document的onmousemove附上不同的事件)
            document.onmousemove = function(mouse_even) {

                if(dragging) {

                    var mouse_even =  mouse_even || windown.even;
                    interval_left = mouse_even.clientX - mouse_down_left;
                    var new_block_left = block_left + interval_left;

                    if(new_block_left < 0 || (new_block_left + block_width + 10) > bar_width) {
                        return false;
                    }
                    // console.log('sum: ' + (block_left + interval_left) + 'bar_width: ' + bar_width + 'block_left:' + block_left + ' interval_left:' + interval_left);
                    the_block.css({
                        "left": new_block_left + "px"
                    });

                    change_cover_height(bgContainerHeight + new_block_left);

                    return false;
                }
            };

            return false;
        });

        $('body').mouseup(function() {
            dragging = false;
        });


    };

    cover_height_drag();


    /**
     * 封面图管理 - 系统提供 - 分页
     * @type {*|jQuery|HTMLElement}
     */

    var sys_support_covers = $('#sys-support-covers');
    var sys_support_pages = $('#cover-sys-support-pages');

    var get_sys_covers = function() {
        var page = 1;
        if(arguments.length > 0) {
            page = arguments[0];
        }

        $.get(user_setting.sys_cover, {page: page}, function(res) {
            res = eval("(" + res + ")");
            if(!res.status) {
                return;
            }

            // 1.1 将封面显示
            sys_support_covers.html('');
            $.each(res.data, function(i, item) {
                // console.log(i + ':' + item.cover_id);
                var item_class = (i % 2) == 0 ? 'cover-item-left' : 'cover-item-right';

                // 1.1.2 append cover_item and 显示选中状态
                var selected_flag = '';
                if(parseInt(new_cover_id_input.val()) == item.cover_id) {
                    selected_flag = '<div class="cover-item-selected">' +
                                        '<img src="' + user_setting.base_url + '/public/img/cover-selected.png">' +
                                    '</div>'
                }
                sys_support_covers.append(
                    '<div class="cover-item img-polaroid ' + item_class + '" cover_id="' + item.cover_id + '" cover_url="'+ item.cover_url +'">' +
                        '<img src="' + item.cover_url + '">' +
                        selected_flag +
                    '</div>');
            });

            // 2.1 绑定每个cover项的点击事件
            // 选择 - 系统提供的封面图
            $('.cover-item').click(function() {
                $('.cover-item-selected').remove();
                $(this).append(
                    '<div class="cover-item-selected">' +
                        '<img src="' + user_setting.base_url + '/public/img/cover-selected.png">' +
                    '</div>'
                );

                var cover_id = parseInt($(this).attr('cover_id'));
                var cover_url = $(this).attr('cover_url');

                new_cover_id_input.val(cover_id);
                new_cover_input.val(cover_url);

                bgStretcher.html('<ul><li><img width="100%" src="' + cover_url + '"></li></ul>');
            });
        });
    };

    // 系统提供 - 获取分页
    sys_support_pages.find('a').click(function() {
        var page_num = parseInt($(this).attr('page'));
        get_sys_covers(page_num);

        sys_support_pages.find('a').attr('class', 'btn');
        $(this).attr('class', 'btn btn-inverse');
    });


    // 保存 - 新封面图
    var cover_save_btn = $('#cover-manage-save');
    cover_save_btn.click(function() {
        $(this).html('保存中...' + '<img width="15" src="' + user_setting.base_url + '/public/img/primary_loading.gif"/>')

        var new_cover_id = parseInt(new_cover_id_input.val());
        var new_cover = new_cover_input.val();
        var new_cover_height = $('#container').height();

        /*if(new_cover_id <= 0 && new_cover_height == parseInt(old_cover_height_input.val())) {
            return;
        }*/

        $.post(user_setting.select_cover, {
            cover_id: isNaN(new_cover_id) ? $('#old-custom-cover-id').val() : new_cover_id,
            cover_height: new_cover_height
         }, function(res) {
             // console.log(res);
            cover_save_btn.html('保存');

             res = eval("(" + res + ")");
             if(!res.status){
                 return;
            }

            // 显示新的封面图
            if(new_cover) { // 如果new_cover为空，表示只修改了高度没有修改封面图
                cover_manage_custom_pic.attr('src', new_cover);
                bgStretcher.html('<ul><li><img width="100%" src="' + new_cover + '"></li></ul>');
                old_custom_cover_input.val(new_cover);
            }

            new_cover_id_input.val('');
            new_cover_input.val('');
            old_cover_height_input.val(new_cover_height);

            layout_cover_manage.fadeOut(100);
        });
    });

    // 调整封面图高度
    var change_cover_height = function(height) {
        $('#container').css({
            height: height + 'px'
        });

        bgStretcher.css({
            height: height + 'px'
        });
    };

    // 设置当前查看的用户封面高度
    change_cover_height(user_setting.cover_height);
    // 管理封面图 - 自定义 - 滑块位置
    the_block.css({
        'left': user_setting.cover_height - user_setting.default_cover_height
    });


    /**
     * 分页的履历
     *
     */
    var exp = {};
    var layout_exp_full = $('#layout-exp-full');
    var layout_exp_date = $('#layout-exp-date');
    var layout_exp_summary = $('#exp-summary');
    var layout_exp_detail = $('#exp-detail');

    exp.next = $('#exp-next');
    exp.container = $('#experience-container');
    exp.close_exp_full = $('#layout-close-exp-full');
    exp.cur_page = 1;
    exp.total_page = user_setting.exp_total_page;

    exp.get_page_data = function() {

        if(exp.cur_page >= exp.total_page) {
            exp.cur_page = exp.total_page;
            exp.next.attr('disabled', 'disabled');
            exp.next.hide();
        }

        exp.next.html('加载中...');
        $.get(user_setting.page_exp, {
            page: exp.cur_page
        }, function(res) {
            res = eval("(" + res + ")");
            // console.log(res.status);
            if(!res.status) {
                return ;
            }

            exp.next.html('加载更多');
            $.each(res.data, function(i, item) {
                exp.container.append(
                    '<div class="exp-item" exp_id="' + item.id + '" inner_link="'+i+'" bg_color="' + item.bg_color + '">' +
                        '<a name="exp-item-'+i+'"></a>' +
                        '<div class="exp-item-header">' +
                            '<span class="exp-item-date-box">' +
                               '<span class="exp-item-date">' + item.start_date + '~' + item.end_date + '</span>' +
                                '<span class="exp-item-date-bg" style="background: ' + item.bg_color + ';"></span>' +
                            '</span>' +
                        '</div>' +

                        '<div class="exp-item-content">' +
                            '<h4>' + item.summary + '</h4>' +
                            '<div class="exp-item-description">' + item.detail + '</div>' +
                            '<div class="exp-item-description-full"></div>' +
                        '</div>' +

                        '<div class="exp-item-footer">' +
                            '<span class="exp-item-full"><a href="javascript:void(0);" class="exp-item-full-link">查看完整</a></span>' +
                        '</div>' +
                    '</div>');
            });

            // 查看完整 履历
            $('.exp-item-full-link').unbind('click').bind('click', function(){
                var curLoadLink =  $(this);
                var curFullDesc = curLoadLink.parents('.exp-item').find('.exp-item-description-full');
                var curOmitDesc = curLoadLink.parents('.exp-item').find('.exp-item-description');
                var exp_id = parseInt(curLoadLink.parents('.exp-item').attr('exp_id'));

                // 收起
                if(curLoadLink.html() == '收起 ↑') {

                    /*curFullDesc.animate({
                        height: curOmitDesc.height()
                    }, 500, function() {
                        curFullDesc.hide();
                        curOmitDesc.show();
                        curLoadLink.html('查看完整');
                    });*/
                    curFullDesc.hide();
                    curOmitDesc.show();
                    curLoadLink.html('查看完整');

                    // 浏览器定位
                    window.location.href = '#exp-item-' + curLoadLink.parents('.exp-item').attr('inner_link');

                    return;
                }else if(curFullDesc.html() != '') {
                    // 如果之前已经记载过，则直接显示
                    // 本来是做成动画下拉，但是发现有的浏览器不兼容，故注释掉
                    /*var omitHeight = curOmitDesc.height();
                    var curFullDescHeight = parseInt(curFullDesc.attr('full_height'));

                    curFullDesc.css({
                        height: omitHeight + 'px'
                    }).show().animate({
                            height: curFullDescHeight
                        });*/
                    curOmitDesc.hide();
                    curFullDesc.show();

                    curLoadLink.html('收起 ↑');
                    return ;
                }

                // 从远程加载
                curLoadLink.html('加载中...');
                $.post(user_setting.exp_view, {
                    exp_id: exp_id
                }, function(res) {
                    res = eval("(" + res + ")");
                    var exp = res.data;

                    curFullDesc.html(exp.detail);
                    /*var omitHeight = curOmitDesc.height();
                    var curFullDescHeight = curFullDesc.height();
                    curFullDesc.attr('full_height', curFullDescHeight);
                    curFullDesc.css({
                        height: omitHeight + 'px'
                    }).show().animate({
                        height: curFullDescHeight
                    });
                    */
                    curOmitDesc.hide();
                    curFullDesc.show();

                    curLoadLink.html('收起 ↑')
                });
            });

            // 日期 - 背景颜色变化
            var over_flag = false;
            $('.exp-item').mouseenter(function() {
                if(!over_flag) {
                    over_flag = true;
                    $(this).find('.exp-item-date-bg').animate({
                        width: '100%'
                    });
                }

                $(this).css({
                    'border-bottom': '1px dashed #ddd'
                });

            });
            $('.exp-item').mouseleave(function() {
                if(over_flag) {
                    over_flag = false;
                    $(this).find('.exp-item-date-bg').animate({
                        width: 5
                    });
                }

                $(this).css({
                    'border-bottom': '1px dashed #eee'
                });

            });

        });
    };


    // 关闭-详细经历-浮层
    exp.close_exp_full.click(function() {
        layout_exp_full.fadeOut(100);

        layout_exp_date.html('');
        layout_exp_detail.html('');
        layout_exp_summary.html('');
    });

    exp.next.click(function() {
        exp.cur_page = exp.cur_page + 1;

        exp.get_page_data();
    });

    // 初始化第一页的履历
    exp.get_page_data();

});
