c = 0
for a in [["1","2"], ["1","3"], ["6","7"], ["5", "7"]]:
    for z in [x for x in range(1,25)]:
        for b in ["","-"]:
            c+=1
            print("| {}_{}_{}{} |".format(a[0],a[1],b,z))
