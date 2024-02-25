import graphviz
g = graphviz.Graph(engine='circo')

def nfold(ls):
    assert(len(ls) > 0)
    if len(ls) == 1: return [[i] for i in ls[0]]
    out = []
    this = ls[0]
    rest = nfold(ls[1:])
    for v in rest:
        for t in this:
            out.append([t] + v)
    return out

def touch(a,b):
    assert len(a) == len(b)
    for i in range(len(a)):
        if a[i] == b[i]: return True
    return False


a = [
        # ["a", "b", "c", "d"],
        ["a", "b", "c"],
        ["a", "b", "c"],
        # ["hazelnut", "walnut", "pistacio", "almond"],
        # ["hazelnut", "walnut"],
        # ["honey", "pekmez"],
]


def mul(l):
    n = 1
    for i in l: n *= i
    return n

def levelsalg(lvls):
    n = mul(lvls)
    prev = [0 for i in lvls]
    out = [prev[:]]
    for i in range(n):
        prev = [(prev[i] + 1) % lvls[i] for i in range(len(prev))]
        out += [prev[:]]
    return out

def maplevels(res, orig):
    return [[b[a] for a,b in zip(r,orig)]for r in res]

ting = levelsalg([len(l) for l in a])
print(ting)
print("here",maplevels(ting, a))





def shorten(l):
    s = ""
    for i in l:
        s += i[0]
    return s

b = nfold(a)

for i in range(len(b)):
    for k in range(i,len(b)):
        o,t = b[i], b[k]
        if not touch(o, t):
            g.edge(" ".join(o), " ".join(t))
print(a)
print(b)

# for i in range(len(a)):
#     for b in range(i, len(a)):
#         g.edge(a[i], a[b])

g.render()
