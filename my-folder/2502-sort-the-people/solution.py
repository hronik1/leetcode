from collections import namedtuple

class Solution:
    def sortPeople(self, names: List[str], heights: List[int]) -> List[str]:
        Person = namedtuple('Person', ['name', 'height'])
        people = [Person(names[i], heights[i]) for i in range(len(names))]
        people.sort(reverse=True, key=attrgetter('height'))
        return [person.name for person in people]
