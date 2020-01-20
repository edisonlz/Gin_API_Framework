--doc see http://wiki.nginx.org/HttpCoreModule#.24uri


function catch(what)
   return what[1]
end

function try(what)
   status, result = pcall(what[1])
   if not status then
      what[2](result)
   end
   return result
end

local ok = nil
local pid = nil
local sig = nil
local doc = nil
local timestamp = 0 

local req_method = ngx.var.request_method
local expire_time = 60 * 60 * 1

local secret = "82406d2ff6c40894a26a3ad34eafff2f"
local args = ngx.req.get_uri_args()

-- Get auth params
if req_method == "GET" then
    pid = ngx.var.arg_pid
    sig = ngx.var.arg__s_
    timestamp = ngx.var.arg__t_
    from = ngx.var.arg_f
    doc = ngx.var.arg_doc
elseif  req_method == "POST" or req_method == "PUT" or  req_method == "DELETE" then
    try {
        function()
            ngx.req.read_body()
            local post_args = ngx.req.get_post_args()
            args = post_args
            pid = post_args.pid
            sig = post_args._s_
            timestamp = post_args._t_
            from = post_args.f
            doc = post_args.doc
        end,
        catch {
              function(error)

              end
        }
    }

    if not sig then
        sig=ngx.var.arg__s_
    end

    if not timestamp then
        timestamp = ngx.var.arg__t_
    end

    if not doc then
        doc = ngx.var.arg_doc
    end


else
    ngx.exit(ngx.HTTP_FORBIDDEN)
end

-- ***** doc debug *****
if doc then
    return 
end

if not timestamp or timestamp == ""  then
    ngx.exit(ngx.HTTP_FORBIDDEN)
end

if not sig or sig == "" then 
    ngx.exit(ngx.HTTP_FORBIDDEN)
end


local filter_args = ""
local key_table = {}  
--取出所有的键  
for key,_ in pairs(args) do  
    table.insert(key_table,key)
end  
--对所有键进行排序  
table.sort(key_table)  
for _,key in pairs(key_table) do
    if key ~= "_s_"  and key ~= "_t_"  then
        filter_args = filter_args .. key .. "=" ..args[key]
    end
end

-- args = 
-- local filter_args, n, err = ngx.re.sub(args, "(&|\\?)_s_=[^&]*&?", "")
-- if filter_args then
-- else
--     filter_args = args
-- end

-- local new_filter_args, nn, eerr = ngx.re.sub(filter_args, "(&|\\?)_t_=[^&]*&?", "")
-- if new_filter_args then
--     filter_args = new_filter_args
-- else
--     if not filter_args then
--         filter_args = args
--     end
-- end

-- ngx.log(ngx.ERR, ngx.var.uri)
-- ngx.log(ngx.ERR, filter_args)

-- the request is expired
-- 从web端过来的请求不校验时间
if from ~= 'w' then
    if ngx.now() - timestamp < -expire_time or ngx.now() - timestamp >= expire_time then
        ngx.status = ngx.HTTP_GONE
        local _now = ngx.now()
        ngx.header.server_time = _now
        ngx.say(_now)
        ngx.exit(ngx.HTTP_OK)
    end
end



-- generate signature
token_string = req_method .. ":" .. ngx.var.uri .. ":" .. filter_args .. ":" .. timestamp .. ":"  .. secret
ngx.log(ngx.ERR, token_string)
token = ngx.md5(token_string)

ngx.log(ngx.ERR, token)

-- Compare sever genned sig(var token) with request sig(request param sig)
if token ~= sig then
    -- ngx.say(token_string)
    -- ngx.say(token)
    ngx.exit(ngx.HTTP_FORBIDDEN)
else
    return
end


