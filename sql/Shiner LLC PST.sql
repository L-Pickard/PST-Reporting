WITH Lead_Time
AS (
    SELECT [Lead Time]
        ,[LT Months]
    FROM (
        VALUES (
             '90DAYS', 3)
           ,('90 DAYS', 3)
           ,('90D', 3)
           ,('3M', 3)
           ,('4M', 4)
           ,('120D', 4)
           ,('60 days', 2)
           ,('60', 2)
           ,('60D', 2)
           ,('2M', 2)
           ,('60', 2)
           ,('3.5M', 3.5)
           ,('4.75M', 4.75)
           ,('4.5M', 4.5)
           ,('6M', 6)
           ,('2.5M', 2.5)
           ,('1.5M', 1.5)
           ,('45 DAYS', 1.5)
           ,('45 days production', 1.5)
           ,('30D', 1)
           ,('1M', 1)
           ,('0.5M', 0.5)
           ,('15 DAYS', 0.5)
           ,('N/A', 6)
           ,('NA', 6)
           ,('0M', 6)
           ,(' ', 6)
           ,('30 days', 1)
		) AS Lead_Time([Lead Time], [LT Months])
	)
	,[Item Images]
AS (
	SELECT di.[Item No]
		,TRIM(di.[Image URL]) AS [Img URL]
		,CASE 
			WHEN ISNUMERIC(SUBSTRING(di.[File Path], CHARINDEX('.', di.[File Path]) - 2, 1)) = 1
				AND ISNUMERIC(SUBSTRING(di.[File Path], CHARINDEX('.', di.[File Path]) - 1, 1)) = 1
				THEN 0
			WHEN ISNUMERIC(SUBSTRING(di.[File Path], CHARINDEX('.', di.[File Path]) - 2, 1)) = 0
				AND ISNUMERIC(SUBSTRING(di.[File Path], CHARINDEX('.', di.[File Path]) - 1, 1)) = 1
				THEN TRY_CONVERT(INTEGER, SUBSTRING(di.[File Path], CHARINDEX('.', di.[File Path]) - 
							1, 1))
			ELSE 100
			END AS [Index]
	
	FROM [dImage] AS di
	
	WHERE LEN(di.[File Path]) > 5
	)
	,[Required Images]
AS (
	SELECT DISTINCT it.[Item No]
		,(
			SELECT TOP 1 [Img URL]
			
			FROM [Item Images]
			
			WHERE [Item No] = it.[Item No]
			
			ORDER BY [Index] ASC
			) AS [Image Path]
	
	FROM [Item Images] AS it
	)

SELECT it.[Item No]
	,it.[Vendor Reference] AS [Vendor Ref]
	,it.[Common Item No] AS [Common Item]
	,br.[Buying Category] AS [Buying Cat]
	,br.[Brand Name]
	,it.[Description]
	,it.[Description 2]
	,it.[Colours]
	,it.[Size 1]
	,it.[Size 1 Unit] AS [Unit]
	,it.[UOM]
	,it.[Category Code] AS [Category]
	,it.[Group Code] AS [Group]
	,it.[Season]
	,it.[Item Info]
	,it.[BC Vendor No] AS [Vendor No]
	,it.[Vendor Name]
	,ri.[Image Path] AS [Image Link]
	,it.[USD Trade]
	,it.[USD SRP]
	,ISNULL(lt.[LT Months], 6) AS [LT Months]
	,it.[On Sale] AS [Pref Sale]
	,it.[Bread & Butter]
	,ISNULL(so.[Shipped 331 to 360 Days Ago], 0) AS [12]
	,ISNULL(so.[Shipped 301 to 330 Days Ago], 0) AS [11]
	,ISNULL(so.[Shipped 271 to 300 Days Ago], 0) AS [10]
	,ISNULL(so.[Shipped 241 to 270 Days Ago], 0) AS [9]
	,ISNULL(so.[Shipped 211 to 240 Days Ago], 0) AS [8]
	,ISNULL(so.[Shipped 181 to 210 Days Ago], 0) AS [7]
	,ISNULL(so.[Shipped 151 to 180 Days Ago], 0) AS [6]
	,ISNULL(so.[Shipped 121 to 150 Days Ago], 0) AS [5]
	,ISNULL(so.[Shipped 91 to 120 Days Ago], 0) AS [4]
	,ISNULL(so.[Shipped 61 to 90 Days Ago], 0) AS [3]
	,ISNULL(so.[Shipped 31 to 60 Days Ago], 0) AS [2]
	,ISNULL(so.[Shipped 1 to 30 Days Ago], 0) AS [1]
	,ISNULL(so.[Shipped in Last 360 Days], 0) AS [L360D]
	,ISNULL(so.[Shipped 30 Day Avg], 0) AS [12M 30D Avg]
	,ISNULL(so.[Shipped in Last 180 Days], 0) AS [L180D]
	,ISNULL(so.[Shipped 30 Day Avg 6M], 0) AS [6M 30D Avg]
	,ISNULL(ui.[Inventory], 0) AS [LLC Inv]
	,ISNULL(po.[PO Qty], 0) AS [PO Qty]
	,ISNULL(sq.[SO Qty], 0) AS [SO Qty]
	,CASE 
		WHEN ((ISNULL(ui.[Inventory], 0) + ISNULL(po.[PO Qty], 0)) - ISNULL(sq.[SO Qty]
				, 0)) / NULLIF(ISNULL(so.[Shipped 30 Day Avg], 0), 0) < ISNULL(lt.[LT Months], 6)
			THEN CAST(1 AS BIT)
		ELSE CAST(0 AS BIT)
		END AS [12M Order]
	,((ISNULL(ui.[Inventory], 0) + ISNULL(po.[PO Qty], 0)) - ISNULL(sq.[SO Qty], 0)) / NULLIF(
		ISNULL(so.[Shipped 30 Day Avg], 0), 0) AS [12M PST]
	,CASE 
		WHEN ((ISNULL(ui.[Inventory], 0) + ISNULL(po.[PO Qty], 0)) - ISNULL(sq.[SO Qty]
				, 0)) / NULLIF(ISNULL(so.[Shipped 30 Day Avg 6M], 0), 0) < ISNULL(lt.[LT Months], 6)
			THEN CAST(1 AS BIT)
		ELSE CAST(0 AS BIT)
		END AS [6M Order]
	,((ISNULL(ui.[Inventory], 0) + ISNULL(po.[PO Qty], 0)) - ISNULL(sq.[SO Qty], 0)) / NULLIF(
		ISNULL(so.[Shipped 30 Day Avg 6M], 0), 0) AS [6M PST]

FROM [dItem] AS it

LEFT JOIN [dBrand] AS br
	ON it.[Brand Code] = br.[Brand Code]

LEFT JOIN Lead_Time AS lt
	ON it.[Lead Time] = lt.[Lead Time]

LEFT JOIN [fShipped PST OG] AS so
	ON it.[Item No] = so.[Item No]
		AND so.[Entity] = 'Shiner LLC'

LEFT JOIN (
	SELECT [Item No]
		,CAST([Inventory] AS DECIMAL(38, 20)) AS [Inventory]
	
	FROM [dInventory]
	
	WHERE [Entity] = 'Shiner LLC'
	) AS ui
	ON it.[Item No] = ui.[Item No]

LEFT JOIN (
	SELECT [Item No]
		,CAST(SUM([Outstanding Qty]) AS DECIMAL(38, 20)) AS [PO Qty]
	
	FROM [Finance].[dbo].[fPurchases]
	
	WHERE [Exclusion] = 0
		AND [Intercompany] = 0
		AND [Entity] = 'Shiner LLC'
	
	GROUP BY [Item No]
	) AS po
	ON it.[Item No] = po.[Item No]

LEFT JOIN (
	SELECT [Item No]
		,CAST(SUM([Outstanding Qty]) AS DECIMAL(38, 20)) AS [SO Qty]
	
	FROM [fOrderbook]
	
	WHERE [Exclusion] = 0
		AND [Intercompany] = 0
		AND [Entity] = 'Shiner LLC'
	
	GROUP BY [Item No]
	) AS sq
	ON it.[Item No] = sq.[Item No]

LEFT JOIN [Required Images] AS ri
	ON it.[Item No] = ri.[Item No]

WHERE it.[Item No] IN (
		SELECT DISTINCT fs.[Item No]
		
		FROM [fSales] AS fs
		
		LEFT JOIN [dItem] AS it
			ON fs.[Item No] = it.[Item No]
		
		WHERE it.[LLC Blocked] = 0
			AND fs.[Entity] = 'Shiner LLC'
		
		UNION ALL
		
		SELECT [Item No]
		
		FROM [dInventory]
		
		WHERE [Entity] = 'Shiner LLC'
		)
	AND it.[Item No] <> '';