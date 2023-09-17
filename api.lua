local cartridge = require('cartridge')
local crud = require('crud')
local uuid = require('uuid')
local json = require('json')

function add_user(request)
    local fullname = request:post_param("fullname")
    local result, err = crud.insert_object('users', { user_id = uuid.new(), fullname = fullname })
    if err ~= nil then
        return { body = json.encode({status = "Error!", error = err}), status = 500 }
    end

    return { body = json.encode({status = "Success!", result = result}), status = 200 }
end

function add_video(request)
    local description = request:post_param("description")
    local result, err = crud.insert_object('videos', { video_id = uuid.new(), description = description, likes = 0 })
    if err ~= nil then
        return { body = json.encode({status = "Error!", error = err}), status = 500 }
    end

    return { body = json.encode({status = "Success!", result = result}), status = 200 }
end

function like_video(request)
    local video_id = request:post_param("video_id")
    local user_id = request:post_param("user_id")

    local result, err = crud.update('videos', uuid.fromstr(video_id), {{'+', 'likes', 1}})
    if err ~= nil then
        return { body = json.encode({status = "Error!", error = err}), status = 500 }
    end

    result, err = crud.insert_object('likes', { like_id = uuid.new(),
                                                video_id = uuid.fromstr(video_id),
                                                user_id = uuid.fromstr(user_id)})
    if err ~= nil then
        return { body = json.encode({status = "Error!", error = err}), status = 500 }
    end

    return { body = json.encode({status = "Success!", result = result}), status = 200 }
end

return {
    add_user = add_user,
    add_video = add_video,
    like_video = like_video,
}