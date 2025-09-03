/*===============================================================================================================================================
Project: Shiner Group Potential Sell Through
Language: T-SQL
Author: Leo Pickard
Version: 1.0
Date: 15/08/2024
=================================================================================================================================================
This query returns the dataset needed to create a new potential sell through xlsx report, it is called during execution of the python script.
=================================================================================================================================================*/

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
    ,Ltd_Sales_Orders
AS (
    SELECT [Item No]
        ,SUM([Outstanding Qty]) AS [SO Qty]
    FROM [fOrderbook]
    WHERE [Customer No] NOT IN (
              'CU110036' -- Shiner LLC (Management Recharge)
            , 'CU110077' -- Shiner LLC
            )
        AND [Entity] = 'Shiner Ltd'
    GROUP BY [Item No]
    )
    ,BV_Sales_Orders
AS (
    SELECT [Item No]
        ,SUM([Outstanding Qty]) AS [SO Qty]
    FROM [fOrderbook]
    WHERE [Customer No] NOT IN (
              'CU110036' -- Shiner LLC (Management Recharge)
            , 'CU110077' -- Shiner LLC
            , 'CU103500' -- Shiner Limited
            )
        AND [Document No] NOT IN (
            SELECT [Document No]
            FROM [fOrderbook]
            WHERE [Entity] = 'Shiner Ltd'
                AND [Document No] IS NOT NULL
            )
        AND [Entity] = 'Shiner B.V'
    GROUP BY [Item No]
    )
    ,LLC_Sales_Orders
AS (
    SELECT [Item No]
        ,SUM([Outstanding Qty]) AS [SO Qty]
    FROM [fOrderbook]
    WHERE [Intercompany] = 0
        AND [Exclusion] = 0
        AND [Entity] = 'Shiner LLC'
    GROUP BY [Item No]
    )
    ,Ltd_LLC_Purchase
AS (
    SELECT [Entity]
        ,[Item No]
        ,SUM([Outstanding Qty]) AS [PO Qty]
    FROM [fPurchases]
    WHERE [Entity] IN ('Shiner Ltd', 'Shiner LLC')
    GROUP BY [Entity]
        ,[Item No]
    )
    ,BV_Purchase
AS (
    SELECT [Item No]
        ,SUM([Outstanding Qty]) AS [PO Qty]
    FROM [fPurchases]
    WHERE [Vendor No] <> 'VE100927' -- Shiner Ltd - Back to Back
        AND [Entity] = 'Shiner B.V'
    GROUP BY [Item No]
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
			) AS [Image Link]
	
	FROM [Item Images] AS it
	)
SELECT it.[Item No]
    ,it.[Vendor Reference]
    ,it.[Common Item No] AS [Parent No]
    ,br.[Buying Category]
    ,br.[Budget Category]
    ,br.[Status] AS [Brand Status]
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
    ,it.[Nav Vendor No]
    ,it.[BC Vendor No]
    ,it.[Vendor Name]
    ,ri.[Image Link]
    ,it.[On Sale] AS [Pref Sale]
    ,it.[Hot Product]
    ,it.[Bread & Butter]
    ,it.[Ltd Blocked]
    ,it.[B.V Blocked]
    ,it.[LLC Blocked]
    ,it.[D2C Web Item]
    ,it.[Owtanet Export]
    ,it.[Web Item]
    ,it.[Ltd GBP Unit Cost] AS [GBP Cost]
    ,it.[GBP Trade] AS [GBP Trade]
    ,it.[GBP SRP] AS [GBP SRP]
    ,it.[B.V EUR Unit Cost] AS [EUR Cost]
    ,it.[EUR Trade] AS [EUR Trade]
    ,it.[EUR SRP] AS [EUR SRP]
    ,it.[LLC USD Unit Cost] AS [USD Cost]
    ,it.[USD Trade] AS [USD Trade]
    ,it.[USD SRP] AS [USD SRP]
    ,ISNULL(lt.[LT Months], 6) AS [LT Months]
    -- Shiner Group Values
    ,ISNULL(gs.[Grp 12], 0) AS [Grp 12]
    ,ISNULL(gs.[Grp 11], 0) AS [Grp 11]
    ,ISNULL(gs.[Grp 10], 0) AS [Grp 10]
    ,ISNULL(gs.[Grp 9], 0) AS [Grp 9]
    ,ISNULL(gs.[Grp 8], 0) AS [Grp 8]
    ,ISNULL(gs.[Grp 7], 0) AS [Grp 7]
    ,ISNULL(gs.[Grp 6], 0) AS [Grp 6]
    ,ISNULL(gs.[Grp 5], 0) AS [Grp 5]
    ,ISNULL(gs.[Grp 4], 0) AS [Grp 4]
    ,ISNULL(gs.[Grp 3], 0) AS [Grp 3]
    ,ISNULL(gs.[Grp 2], 0) AS [Grp 2]
    ,ISNULL(gs.[Grp 1], 0) AS [Grp 1]
    ,ISNULL(gs.[Grp L360D], 0) AS [Grp L360D]
    ,ISNULL(gs.[Grp L180D], 0) AS [Grp L180D]
    ,ISNULL((CAST(gs.[Grp L360D] AS DECIMAL(20, 8)) / 12.0), 0.0) AS [Grp 12M 30D Avg]
    ,ISNULL((CAST(gs.[Grp L180D] AS DECIMAL(20, 8)) / 6.0), 0.0) AS [Grp 6M 30D Avg]
    ,ISNULL(gi.[Grp Free Stock], 0) AS [Grp Free Stock]
    ,ISNULL(gi.[Grp Inventory], 0) AS [Grp Inventory]
    ,ISNULL(gb.[Grp Buffer], 0) AS [Grp Buffer]
    ,ISNULL(gp.[Grp PO], 0) AS [Grp PO]
    ,ISNULL(gr.[Grp SO], 0) AS [Grp SO]
    ,CAST(CASE 
            WHEN CAST((
                        (ISNULL(gi.[Grp Inventory], 0.0) + ISNULL(gp.[Grp PO], 0.0)
                            ) - ISNULL(gr.[Grp SO], 0)
                        ) AS DECIMAL(20, 8)) / NULLIF(ISNULL((CAST(gs.[Grp L360D] AS DECIMAL(20, 8)) / 12.0
                            ), 0.0), 0.0) < ISNULL(lt.[LT Months], 6)
                THEN 1
            ELSE 0
            END AS BIT) AS [Grp 12M Order]
    ,CAST((
            (ISNULL(gi.[Grp Inventory], 0.0) + ISNULL(gp.[Grp PO], 0.0)) - ISNULL(gr.
                [Grp SO], 0)
            ) AS DECIMAL(20, 8)) / NULLIF(ISNULL((CAST(gs.[Grp L360D] AS DECIMAL(20, 8)) / 12.0
                ), 0.0), 0.0) AS [Grp 12M PST]
    ,CAST(CASE 
            WHEN CAST((
                        (ISNULL(gi.[Grp Inventory], 0.0) + ISNULL(gp.[Grp PO], 0.0)
                            ) - ISNULL(gr.[Grp SO], 0)
                        ) AS DECIMAL(20, 8)) / NULLIF(ISNULL((CAST(gs.[Grp L180D] AS DECIMAL(20, 8)) / 6.0
                            ), 0.0), 0.0) < ISNULL(lt.[LT Months], 6)
                THEN 1
            ELSE 0
            END AS BIT) AS [Grp 6M Order]
    ,CAST((
            (ISNULL(gi.[Grp Inventory], 0.0) + ISNULL(gp.[Grp PO], 0.0)) - ISNULL(gr.
                [Grp SO], 0)
            ) AS DECIMAL(20, 8)) / NULLIF(ISNULL((CAST(gs.[Grp L180D] AS DECIMAL(20, 8)) / 6.0
                ), 0.0), 0.0) AS [Grp 6M PST]
    -- Shiner Ltd Values
    ,ISNULL(ls.[Shipped 331 to 360 Days Ago], 0) AS [Ltd 12]
    ,ISNULL(ls.[Shipped 301 to 330 Days Ago], 0) AS [Ltd 11]
    ,ISNULL(ls.[Shipped 271 to 300 Days Ago], 0) AS [Ltd 10]
    ,ISNULL(ls.[Shipped 241 to 270 Days Ago], 0) AS [Ltd 9]
    ,ISNULL(ls.[Shipped 211 to 240 Days Ago], 0) AS [Ltd 8]
    ,ISNULL(ls.[Shipped 181 to 210 Days Ago], 0) AS [Ltd 7]
    ,ISNULL(ls.[Shipped 151 to 180 Days Ago], 0) AS [Ltd 6]
    ,ISNULL(ls.[Shipped 121 to 150 Days Ago], 0) AS [Ltd 5]
    ,ISNULL(ls.[Shipped 91 to 120 Days Ago], 0) AS [Ltd 4]
    ,ISNULL(ls.[Shipped 61 to 90 Days Ago], 0) AS [Ltd 3]
    ,ISNULL(ls.[Shipped 31 to 60 Days Ago], 0) AS [Ltd 2]
    ,ISNULL(ls.[Shipped 1 to 30 Days Ago], 0) AS [Ltd 1]
    ,ISNULL(ls.[Shipped in Last 360 Days], 0) AS [Ltd L360D]
    ,ISNULL(ls.[Shipped in Last 180 Days], 0) AS [Ltd L180D]
    ,ISNULL((CAST(ls.[Shipped in Last 360 Days] AS DECIMAL(20, 8)) / 12.0), 0.0) AS 
    [Ltd 12M 30D Avg]
    ,ISNULL((CAST(ls.[Shipped in Last 180 Days] AS DECIMAL(20, 8)) / 6.0), 0.0) AS 
    [Ltd 6M 30D Avg]
    ,ISNULL(li.[Free Stock], 0) AS [Ltd Free Stock]
    ,ISNULL(li.[Inventory], 0) AS [Ltd Inventory]
    ,ISNULL(it.[Ltd Buffer Stock], 0) AS [Ltd Buffer]
    ,ISNULL(lp.[PO Qty], 0) AS [Ltd PO]
    ,ISNULL(lo.[SO Qty], 0) AS [Ltd SO]
    ,CAST(CASE 
            WHEN CAST((
                        (ISNULL(li.[Inventory], 0.0) + ISNULL(lp.[PO Qty], 0.0)
                            ) - ISNULL(lo.[SO Qty], 0)
                        ) AS DECIMAL(20, 8)) / NULLIF(ISNULL((CAST(ls.[Shipped in Last 360 Days] AS DECIMAL(20, 8)) / 12.0
                            ), 0.0), 0.0) < ISNULL(lt.[LT Months], 6)
                THEN 1
            ELSE 0
            END AS BIT) AS [Ltd 12M Order]
    ,CAST((
            (ISNULL(li.[Inventory], 0.0) + ISNULL(lp.[PO Qty], 0.0)) - ISNULL(lo.
                [SO Qty], 0)
            ) AS DECIMAL(20, 8)) / NULLIF(ISNULL((CAST(ls.[Shipped in Last 360 Days] AS DECIMAL(20, 8)) / 12.0
                ), 0.0), 0.0) AS [Ltd 12M PST]
    ,CAST(CASE 
            WHEN CAST((
                        (ISNULL(li.[Inventory], 0.0) + ISNULL(lp.[PO Qty], 0.0)
                            ) - ISNULL(lo.[SO Qty], 0)
                        ) AS DECIMAL(20, 8)) / NULLIF(ISNULL((CAST(ls.[Shipped in Last 180 Days] AS DECIMAL(20, 8)) / 6.0
                            ), 0.0), 0.0) < ISNULL(lt.[LT Months], 6)
                THEN 1
            ELSE 0
            END AS BIT) AS [Ltd 6M Order]
    ,CAST((
            (ISNULL(li.[Inventory], 0.0) + ISNULL(lp.[PO Qty], 0.0)) - ISNULL(lo.
                [SO Qty], 0)
            ) AS DECIMAL(20, 8)) / NULLIF(ISNULL((CAST(ls.[Shipped in Last 180 Days] AS DECIMAL(20, 8)) / 6.0
                ), 0.0), 0.0) AS [Ltd 6M PST]
    -- Shiner B.V Values
    ,ISNULL(bs.[Shipped 331 to 360 Days Ago], 0) AS [B.V 12]
    ,ISNULL(bs.[Shipped 301 to 330 Days Ago], 0) AS [B.V 11]
    ,ISNULL(bs.[Shipped 271 to 300 Days Ago], 0) AS [B.V 10]
    ,ISNULL(bs.[Shipped 241 to 270 Days Ago], 0) AS [B.V 9]
    ,ISNULL(bs.[Shipped 211 to 240 Days Ago], 0) AS [B.V 8]
    ,ISNULL(bs.[Shipped 181 to 210 Days Ago], 0) AS [B.V 7]
    ,ISNULL(bs.[Shipped 151 to 180 Days Ago], 0) AS [B.V 6]
    ,ISNULL(bs.[Shipped 121 to 150 Days Ago], 0) AS [B.V 5]
    ,ISNULL(bs.[Shipped 91 to 120 Days Ago], 0) AS [B.V 4]
    ,ISNULL(bs.[Shipped 61 to 90 Days Ago], 0) AS [B.V 3]
    ,ISNULL(bs.[Shipped 31 to 60 Days Ago], 0) AS [B.V 2]
    ,ISNULL(bs.[Shipped 1 to 30 Days Ago], 0) AS [B.V 1]
    ,ISNULL(bs.[Shipped in Last 360 Days], 0) AS [B.V L360D]
    ,ISNULL(bs.[Shipped in Last 180 Days], 0) AS [B.V L180D]
    ,ISNULL((CAST(bs.[Shipped in Last 360 Days] AS DECIMAL(20, 8)) / 12.0), 0.0) AS 
    [B.V 12M 30D Avg]
    ,ISNULL((CAST(bs.[Shipped in Last 180 Days] AS DECIMAL(20, 8)) / 6.0), 0.0) AS 
    [B.V 6M 30D Avg]
    ,ISNULL(bi.[Free Stock], 0) AS [B.V Free Stock]
    ,ISNULL(bi.[Inventory], 0) AS [B.V Inventory]
    ,ISNULL(it.[B.V Buffer Stock], 0) AS [B.V Buffer]
    ,ISNULL(bp.[PO Qty], 0) AS [B.V PO]
    ,ISNULL(bo.[SO Qty], 0) AS [B.V SO]
    ,CAST(CASE 
            WHEN CAST((
                        (ISNULL(bi.[Inventory], 0.0) + ISNULL(bp.[PO Qty], 0.0)
                            ) - ISNULL(bo.[SO Qty], 0)
                        ) AS DECIMAL(20, 8)) / NULLIF(ISNULL((CAST(bs.[Shipped in Last 360 Days] AS DECIMAL(20, 8)) / 12.0
                            ), 0.0), 0.0) < ISNULL(lt.[LT Months], 6)
                THEN 1
            ELSE 0
            END AS BIT) AS [B.V 12M Order]
    ,CAST((
            (ISNULL(bi.[Inventory], 0.0) + ISNULL(bp.[PO Qty], 0.0)) - ISNULL(bo.
                [SO Qty], 0)
            ) AS DECIMAL(20, 8)) / NULLIF(ISNULL((CAST(bs.[Shipped in Last 360 Days] AS DECIMAL(20, 8)) / 12.0
                ), 0.0), 0.0) AS [B.V 12M PST]
    ,CAST(CASE 
            WHEN CAST((
                        (ISNULL(bi.[Inventory], 0.0) + ISNULL(bp.[PO Qty], 0.0)
                            ) - ISNULL(bo.[SO Qty], 0)
                        ) AS DECIMAL(20, 8)) / NULLIF(ISNULL((CAST(bs.[Shipped in Last 180 Days] AS DECIMAL(20, 8)) / 6.0
                            ), 0.0), 0.0) < ISNULL(lt.[LT Months], 6)
                THEN 1
            ELSE 0
            END AS BIT) AS [B.V 6M Order]
    ,CAST((
            (ISNULL(bi.[Inventory], 0.0) + ISNULL(bp.[PO Qty], 0.0)) - ISNULL(bo.
                [SO Qty], 0)
            ) AS DECIMAL(20, 8)) / NULLIF(ISNULL((CAST(bs.[Shipped in Last 180 Days] AS DECIMAL(20, 8)) / 6.0
                ), 0.0), 0.0) AS [B.V 6M PST]
    -- Shiner LLC Values
    ,ISNULL(us.[Shipped 331 to 360 Days Ago], 0) AS [LLC 12]
    ,ISNULL(us.[Shipped 301 to 330 Days Ago], 0) AS [LLC 11]
    ,ISNULL(us.[Shipped 271 to 300 Days Ago], 0) AS [LLC 10]
    ,ISNULL(us.[Shipped 241 to 270 Days Ago], 0) AS [LLC 9]
    ,ISNULL(us.[Shipped 211 to 240 Days Ago], 0) AS [LLC 8]
    ,ISNULL(us.[Shipped 181 to 210 Days Ago], 0) AS [LLC 7]
    ,ISNULL(us.[Shipped 151 to 180 Days Ago], 0) AS [LLC 6]
    ,ISNULL(us.[Shipped 121 to 150 Days Ago], 0) AS [LLC 5]
    ,ISNULL(us.[Shipped 91 to 120 Days Ago], 0) AS [LLC 4]
    ,ISNULL(us.[Shipped 61 to 90 Days Ago], 0) AS [LLC 3]
    ,ISNULL(us.[Shipped 31 to 60 Days Ago], 0) AS [LLC 2]
    ,ISNULL(us.[Shipped 1 to 30 Days Ago], 0) AS [LLC 1]
    ,ISNULL(us.[Shipped in Last 360 Days], 0) AS [LLC L360D]
    ,ISNULL(us.[Shipped in Last 180 Days], 0) AS [LLC L180D]
    ,ISNULL((CAST(us.[Shipped in Last 360 Days] AS DECIMAL(20, 8)) / 12.0), 0.0) AS 
    [LLC 12M 30D Avg]
    ,ISNULL((CAST(us.[Shipped in Last 180 Days] AS DECIMAL(20, 8)) / 6.0), 0.0) AS 
    [LLC 6M 30D Avg]
    ,ISNULL(ui.[Free Stock], 0) AS [LLC Free Stock]
    ,ISNULL(ui.[Inventory], 0) AS [LLC Inventory]
    ,ISNULL(it.[LLC Buffer Stock], 0) AS [LLC Buffer]
    ,ISNULL(up.[PO Qty], 0) AS [LLC PO]
    ,ISNULL(uo.[SO Qty], 0) AS [LLC SO]
    ,CAST(CASE 
            WHEN CAST((
                        (ISNULL(ui.[Inventory], 0.0) + ISNULL(up.[PO Qty], 0.0)
                            ) - ISNULL(uo.[SO Qty], 0)
                        ) AS DECIMAL(20, 8)) / NULLIF(ISNULL((CAST(us.[Shipped in Last 360 Days] AS DECIMAL(20, 8)) / 12.0
                            ), 0.0), 0.0) < ISNULL(lt.[LT Months], 6)
                THEN 1
            ELSE 0
            END AS BIT) AS [LLC 12M Order]
    ,CAST((
            (ISNULL(ui.[Inventory], 0.0) + ISNULL(up.[PO Qty], 0.0)) - ISNULL(uo.
                [SO Qty], 0)
            ) AS DECIMAL(20, 8)) / NULLIF(ISNULL((CAST(us.[Shipped in Last 360 Days] AS DECIMAL(20, 8)) / 12.0
                ), 0.0), 0.0) AS [LLC 12M PST]
    ,CAST(CASE 
            WHEN CAST((
                        (ISNULL(ui.[Inventory], 0.0) + ISNULL(up.[PO Qty], 0.0)
                            ) - ISNULL(uo.[SO Qty], 0)
                        ) AS DECIMAL(20, 8)) / NULLIF(ISNULL((CAST(us.[Shipped in Last 180 Days] AS DECIMAL(20, 8)) / 6.0
                            ), 0.0), 0.0) < ISNULL(lt.[LT Months], 6)
                THEN 1
            ELSE 0
            END AS BIT) AS [LLC 6M Order]
    ,CAST((
            (ISNULL(ui.[Inventory], 0.0) + ISNULL(up.[PO Qty], 0.0)) - ISNULL(uo.
                [SO Qty], 0)
            ) AS DECIMAL(20, 8)) / NULLIF(ISNULL((CAST(us.[Shipped in Last 180 Days] AS DECIMAL(20, 8)) / 6.0
                ), 0.0), 0.0) AS [LLC 6M PST]
FROM [dItem] AS it
LEFT JOIN [dBrand] AS br
    ON it.[Brand Code] = br.[Brand Code]
LEFT JOIN Lead_Time AS lt
    ON it.[Lead Time] = lt.[Lead Time]
LEFT JOIN (
    SELECT [Item No]
        ,ISNULL([Ltd Buffer Stock], 0) + ISNULL([B.V Buffer Stock], 0) + ISNULL([LLC Buffer Stock], 0) 
        AS [Grp Buffer]
    FROM [dItem]
    ) AS gb
    ON it.[Item No] = gb.[Item No]
-- Shipped qty joins
LEFT JOIN (
    SELECT [Item No]
        ,SUM([Shipped 331 to 360 Days Ago]) AS [Grp 12]
        ,SUM([Shipped 301 to 330 Days Ago]) AS [Grp 11]
        ,SUM([Shipped 271 to 300 Days Ago]) AS [Grp 10]
        ,SUM([Shipped 241 to 270 Days Ago]) AS [Grp 9]
        ,SUM([Shipped 211 to 240 Days Ago]) AS [Grp 8]
        ,SUM([Shipped 181 to 210 Days Ago]) AS [Grp 7]
        ,SUM([Shipped 151 to 180 Days Ago]) AS [Grp 6]
        ,SUM([Shipped 121 to 150 Days Ago]) AS [Grp 5]
        ,SUM([Shipped 91 to 120 Days Ago]) AS [Grp 4]
        ,SUM([Shipped 61 to 90 Days Ago]) AS [Grp 3]
        ,SUM([Shipped 31 to 60 Days Ago]) AS [Grp 2]
        ,SUM([Shipped 1 to 30 Days Ago]) AS [Grp 1]
        ,SUM([Shipped in Last 360 Days]) AS [Grp L360D]
        ,SUM([Shipped in Last 180 Days]) AS [Grp L180D]
    FROM [fShipped Qty]
    GROUP BY [Item No]
    ) AS gs
    ON it.[Item No] = gs.[Item No]
LEFT JOIN [fShipped Qty] AS ls
    ON it.[Item No] = ls.[Item No]
        AND ls.[Entity] = 'Shiner Ltd'
LEFT JOIN [fShipped Qty] AS bs
    ON it.[Item No] = bs.[Item No]
        AND bs.[Entity] = 'Shiner B.V'
LEFT JOIN [fShipped Qty] AS us
    ON it.[Item No] = us.[Item No]
        AND us.[Entity] = 'Shiner LLC'
-- Purchase qty joins
LEFT JOIN (
    SELECT it.[Item No]
        ,ISNULL(ll.[PO Qty], 0) + ISNULL(pb.[PO Qty], 0) AS [Grp PO]
    FROM [dItem] AS it
    LEFT JOIN (
        SELECT [Item No]
            ,SUM([PO Qty]) AS [PO Qty]
        FROM Ltd_LLC_Purchase
        GROUP BY [Item No]
        ) AS ll
        ON it.[Item No] = ll.[Item No]
    LEFT JOIN BV_Purchase AS pb
        ON it.[Item No] = pb.[Item No]
    WHERE ISNULL(ll.[PO Qty], 0) + ISNULL(pb.[PO Qty], 0) <> 0
    ) AS gp
    ON it.[Item No] = gp.[Item No]
LEFT JOIN Ltd_LLC_Purchase AS lp
    ON it.[Item No] = lp.[Item No]
        AND lp.[Entity] = 'Shiner Ltd'
LEFT JOIN BV_Purchase AS bp
    ON it.[Item No] = bp.[Item No]
LEFT JOIN Ltd_LLC_Purchase AS up
    ON it.[Item No] = up.[Item No]
        AND up.[Entity] = 'Shiner LLC'
-- Sales order qty joins
LEFT JOIN (
    SELECT it.[Item No]
        ,ISNULL(lo.[SO Qty], 0) + ISNULL(bo.[SO Qty], 0) + ISNULL(uo.[SO Qty], 0) AS [Grp SO]
    FROM [dItem] AS it
    LEFT JOIN Ltd_Sales_Orders AS lo
        ON it.[Item No] = lo.[Item No]
    LEFT JOIN BV_Sales_Orders AS bo
        ON it.[Item No] = bo.[Item No]
    LEFT JOIN LLC_Sales_Orders AS uo
        ON it.[Item No] = uo.[Item No]
    WHERE ISNULL(lo.[SO Qty], 0) + ISNULL(bo.[SO Qty], 0) + ISNULL(uo.[SO Qty], 0) <> 0
    ) AS gr
    ON it.[Item No] = gr.[Item No]
LEFT JOIN Ltd_Sales_Orders AS lo
    ON it.[Item No] = lo.[Item No]
LEFT JOIN BV_Sales_Orders AS bo
    ON it.[Item No] = bo.[Item No]
LEFT JOIN LLC_Sales_Orders AS uo
    ON it.[Item No] = uo.[Item No]
-- Inventory Qty Joins
LEFT JOIN (
    SELECT [Item No]
        ,SUM([Free Stock]) AS [Grp Free Stock]
        ,SUM([Inventory]) AS [Grp Inventory]
    FROM [dInventory]
    GROUP BY [Item No]
    ) AS gi
    ON it.[Item No] = gi.[Item No]
LEFT JOIN [dInventory] AS li
    ON it.[Item No] = li.[Item No]
        AND li.[Entity] = 'Shiner Ltd'
LEFT JOIN [dInventory] AS bi
    ON it.[Item No] = bi.[Item No]
        AND bi.[Entity] = 'Shiner B.V'
LEFT JOIN [dInventory] AS ui
    ON it.[Item No] = ui.[Item No]
        AND ui.[Entity] = 'Shiner LLC'
LEFT JOIN [Required Images] AS ri
	ON it.[Item No] = ri.[Item No]
WHERE NOT (
        it.[Ltd Blocked] = 1
        AND it.[B.V Blocked] = 1
        AND (
            [BC Vendor No] IS NULL
            OR [LLC Blocked] = 1
            )
        )
    AND it.[Item No] <> '';