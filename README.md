Go Simple Miner
یک ماینر ساده جهت تفهیم مسله ماینینگ در زبان گو
با ده خط کد ماینر خود را داشته باشید

الگورتیم مورد استفاده sha256
امکان افزایش سختی مسله
✨نوشته شده در زبان go ✨
نصب
git clone https://github.com/thisismz/Go_Simple_Miner.git
cd Go_Simple_Miner
go build
اجرا

./Go_Simple_Miner.exe
توضحیات کد
تابع ای که عملکرد ماینینگ را انجام میدهد

func mine(o interface{}, difficulty int) {...

مخشص کردن سختی
هر چه عدد بزرگتر باشد تعداد صفر ها در perfix بیشتر شده وپیدا کردن nounce سختتر خواهد بود

difficulty := 5 // increase number for make mining harder

خط 47 تا 64
transone := SimpleTransaction{ ID: 1, from: "mahdi", to: "mohammad", amount: 25, } tanstwo := SimpleTransaction{ ID: 2, from: "roya", to: "negin", amount: 99, } block := SimpleBlock{ ID: 1, previousBlockHash: "00000891861fd09d3cf161db54434c9e518cf80e94811e7762c3aee8a7af39af", transactions: []SimpleTransaction{transone, tanstwo}, }
مربوط به تعریف تراکنش ها و بلاک هست محدودتی در تعریف تعداد وجود ندارد
