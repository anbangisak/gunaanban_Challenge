import re
ccList = [
    "4123456789123456", "5123-4567-8912-3456", "61234-567-8912-3456", "4123356789123456", "5133-3367-8912-3456",
	"5123 - 3567 - 8912 - 3456",
]
# error: bad character in group name '?<=\\d{4}' at position 17
# python also doesnt support it
reCompileObj = re.search(r'\b((?!^[0-37-9])(?(?<=[0-9]{4})-?|)(\d)(?!\2{2}-\2|\2-\2{2}|-\2{3})){16}\b', ccList[0])
for item in ccList:
    print("{0} : {1}".format(item, reCompileObj.match(item)))
