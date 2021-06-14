1.  Nick works at a clothing store. He has a large pile of socks that he must pair by color for
    sale. Given an array of integers representing the color of each sock, determine how many
    pairs of socks with matching colors there are.
    For example, there are n = 7 socks with colors ar = [1, 2, 1, 2, 1, 3, 2]. There is one pair of
    color 1 and one of color 2. There are three odd socks left, one of each color. The number of
    pairs is 2.Function Description
    Complete the sock merchant function in the editor below. It must return an integer
    representing the number of matching pairs of socks that are available.
    sockMerchant has the following parameter(s):<br />
    ● n : the number of socks in the pile<br />
    ● ar: the colors of each sock<br />
    Input Format<br />
    The first line contains an integer n , the number of socks represented in a r .
    The second line contains n space-separated integers describing the colors ar [ i ] of the
    socks in the pile.<br />
    Constraints :<br />
    1 ≤ n ≤ 100<br />
    1 ≤ a r [ i ] ≤ 100 w here 0 ≤ i < n<br />
    Output Format<br />
    Return the total number of matching pairs of socks that Nick can sell.
    Sample Input : 9<br />
    arr : [10, 20, 20, 10, 10, 30, 50, 10, 20]<br />
    Sample output : 3<br />

2.  Bill is an avid hiker. He tracks his hikes meticulously, paying close attention to small details
    like topography. During his last hike he took exactly n steps. For every step he took, he
    noted if it was an uphill, U , or a downhill, D step. Gary's hikes start and end at sea level
    and each step up or down represents a 1 unit change in altitude. We define the following
    terms:<br />
    ● A mountain is a sequence of consecutive steps above sea level, starting with a step
    up from sea level and ending with a step down to sea level.<br />
    ● A valley is a sequence of consecutive steps below sea level, starting with a step
    down from sea level and ending with a step up to sea level.<br />
    Given Gary's sequence of up and down steps during his last hike, find and print the number
    of valleys he walked through.
    For example, if Gary's path is s = [D D U U U U D D] , he first enters a valley 2 units deep.
    Then he climbs out an up onto a mountain 2 units high. Finally, he returns to sea level and
    ends his hike.Function Description
    Complete the countingValleys function in the editor below. It must return an integer that
    denotes the number of valleys Gary traversed.<br />
    countingValleys has the following parameter(s):<br />
    ● n : the number of steps Gary takes<br />
    ● s : a string describing his path<br />
    Input Format<br />
    The first line contains an integer n , the number of steps in Gary's hike.<br />
    The second line contains a single string s , of n characters that describe his path.<br />
    Constraints<br />
    ● 2 ≤ n ≤ 106<br />
    ● s [i] ∈ { U D }<br />
    Output Format<br />
    Print a single integer that denotes the number of valleys Gary walked through during his
    hike.<br />
    Sample Input : 8<br />
    arr : [U, D, D, D, U, D, U, U]<br />
    Sample Output : 1<br />

3.  There is an input number: 1.345.679<br />
    Write pseudo code (use ​ GoLang​ ) that produces following output:<br />
    1000000<br />
    300000<br />
    40000<br />
    5000<br />
    600<br />
    70<br />
    9<br />

4.  Andrew walks through 100 switches from point A to point B with 1 to 100 as the number.
    At the first trip, Andrew presses all of the switches so lamps are on. Second trip, Andrew
    only presses switches that multiplying of two. The next trip, Andrew presses switches
    that multiplying of three. Andrew repeats his trips from A to B for 100 times. Write down
    the code to determine how many lamps will turn on after 100 trips from A to B.