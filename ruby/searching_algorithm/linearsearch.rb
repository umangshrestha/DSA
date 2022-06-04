def linear_search(arr, target)
    arr.each_with_index do |v, i|
        return i if v == target
    end
    return -1
end


