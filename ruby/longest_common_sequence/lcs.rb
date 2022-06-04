def max(x, y)
    (x>y)? x:y
end

def longest_common_sequence(str1, str2)
    l1 = str1.length
    l2 = str2.length
    a = Array.new(l1+1){Array.new(l2+1) {0} }
    (l1-1).downto(-1) do |i|
        (l2-1).downto(-2) do |j|
            if str1[i] == str2[j]
                a[i][j] = a[i+1][j+1] + 1
            else
                a[i][j] = max(a[i][j+1], a[i+1][j])
            end
        end
    end
    return a[0][0]
end
